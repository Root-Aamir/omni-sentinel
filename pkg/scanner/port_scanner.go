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
	Target    string
	StartPort int
	EndPort   int
	TeleToken string
	TeleID    string
}

func (s Scout) Name() string { return "Network-Scout" }

func (s Scout) Execute() error {
	var wg sync.WaitGroup
	for p := s.StartPort; p <= s.EndPort; p++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", s.Target, port)
			conn, err := net.DialTimeout("tcp", address, 2*time.Second)
			if err == nil {
				// Banner Grabbing Logic
				conn.SetReadDeadline(time.Now().Add(2 * time.Second))
				buffer := make([]byte, 1024)
				n, _ := conn.Read(buffer)
				service := "Unknown Service"
				if n > 0 {
					service = strings.TrimSpace(string(buffer[:n]))
				}

				msg := fmt.Sprintf("🎯 PORT OPEN: %d\n🔍 Service: %s\n🌐 Host: %s", port, service, s.Target)
				fmt.Println("[!]", msg)
				utils.SaveLog("Scout", "FOUND", msg)
				if s.TeleToken != "" {
					utils.SendTelegramAlert(s.TeleToken, s.TeleID, msg)
				}
				conn.Close()
			}
		}(p)
	}
	wg.Wait()
	return nil
}
