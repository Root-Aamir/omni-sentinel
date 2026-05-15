package scanner

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/Root-Aamir/omni-sentinel/pkg/utils"
)

type Scout struct {
	Targets   []string
	StartPort int
	EndPort   int
	TeleToken string
	TeleID    string
}

func (s Scout) Name() string { return "Multi-Target-Scout" }

func (s Scout) Execute() error {
	var wg sync.WaitGroup

	for _, target := range s.Targets {
		fmt.Printf("[🔎] Initiating Scan on: %s\n", target)

		for p := s.StartPort; p <= s.EndPort; p++ {
			wg.Add(1)
			go func(t string, port int) {
				defer wg.Done()
				address := fmt.Sprintf("%s:%d", t, port)
				// Dial timeout ko thoda badhaya hai connectivity ke liye
				conn, err := net.DialTimeout("tcp", address, 3*time.Second)
				if err == nil {
					// Banner Grabbing
					conn.SetReadDeadline(time.Now().Add(2 * time.Second))
					buffer := make([]byte, 1024)
					n, _ := conn.Read(buffer)
					service := "Unknown Service"
					if n > 0 {
						service = strings.TrimSpace(string(buffer[:n]))
					}

					msg := fmt.Sprintf("🎯 PORT OPEN: %d\n🔍 Service: %s\n🌐 Host: %s", port, service, t)
					fmt.Println("[!]", msg)
					utils.SendTelegramAlert(s.TeleToken, s.TeleID, msg)
					conn.Close()
				}
			}(target, p)
			// Bot detection se bachne ke liye chota delay
			time.Sleep(10 * time.Millisecond)
		}
	}
	wg.Wait()
	return nil
}
