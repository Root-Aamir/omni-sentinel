package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type Scout struct {
	Target    string
	StartPort int
	EndPort   int
}

func (s Scout) Name() string { return "Scout (Port Scanner)" }

func (s Scout) Execute() error {
	var wg sync.WaitGroup
	fmt.Printf("[+] Scout: Scanning %s (%d-%d)...\n", s.Target, s.StartPort, s.EndPort)

	for p := s.StartPort; p <= s.EndPort; p++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", s.Target, port)
			conn, err := net.DialTimeout("tcp", address, 1*time.Second)
			if err == nil {
				fmt.Printf("   [!] OPEN: %d\n", port)
				conn.Close()
			}
		}(p)
	}
	wg.Wait()
	return nil
}
