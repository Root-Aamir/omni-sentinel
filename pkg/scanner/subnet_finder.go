package scanner

import (
	"fmt"
	"os/exec"
	"runtime"
	"sync"
)

func PingTarget(ip string) bool {
	var cmd *exec.Cmd
	// Timeout ko thoda badhaya hai taaki slow Wi-Fi par bhi response mil jaye
	if runtime.GOOS == "windows" {
		cmd = exec.Command("ping", "-n", "1", "-w", "800", ip)
	} else {
		cmd = exec.Command("ping", "-c", "1", "-W", "1", ip)
	}
	return cmd.Run() == nil
}

func DiscoverHosts(baseIP string) []string {
	var aliveHosts []string
	var wg sync.WaitGroup
	var mu sync.Mutex

	fmt.Printf("[📡] Scanning Subnet: %s.0/24...\n", baseIP)

	// Scanning 254 IPs concurrently
	for i := 1; i < 255; i++ {
		wg.Add(1)
		targetIP := fmt.Sprintf("%s.%d", baseIP, i)
		go func(ip string) {
			defer wg.Done()
			if PingTarget(ip) {
				mu.Lock()
				aliveHosts = append(aliveHosts, ip)
				mu.Unlock()
				fmt.Printf("    [+] Discovered: %s\n", ip)
			}
		}(targetIP)
	}
	wg.Wait()

	if len(aliveHosts) == 0 {
		fmt.Println("[!] No other active devices found. Check your base IP or Firewall.")
	}
	return aliveHosts
}
