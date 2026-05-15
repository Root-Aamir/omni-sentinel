package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/Root-Aamir/omni-sentinel/pkg/utils" // Logger utility import ki
)

// Scout structure engine ke Task interface ko satisfy karti hai
type Scout struct {
	Target    string
	StartPort int
	EndPort   int
}

// Name module ka pehchan batata hai
func (s Scout) Name() string {
	return "Network-Scout (Port Scanner)"
}

// Execute mein asli scanning logic aur logging hai
func (s Scout) Execute() error {
	var wg sync.WaitGroup
	fmt.Printf("[+] Scout: Initiating high-speed scan on %s...\n", s.Target)

	// Concurrency: Har port ke liye ek alag Goroutine
	for port := s.StartPort; port <= s.EndPort; port++ {
		wg.Add(1)

		go func(p int) {
			defer wg.Done()

			address := fmt.Sprintf("%s:%d", s.Target, p)
			// 1 second timeout for professional speed
			conn, err := net.DialTimeout("tcp", address, 1*time.Second)

			if err == nil {
				// Agar port open hai
				resultMsg := fmt.Sprintf("Port %d is OPEN", p)
				fmt.Printf("   [!] %s\n", resultMsg)

				// JSON Log save karein (Persistence)
				utils.SaveLog("Scout", "VULN-INFO", map[string]interface{}{
					"target": s.Target,
					"port":   p,
					"status": "open",
				})

				conn.Close()
			}
		}(port)
	}

	wg.Wait()
	fmt.Printf("[✔] Scout: Scan on %s completed.\n", s.Target)
	return nil
}
