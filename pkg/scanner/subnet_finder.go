package scanner

import (
	"fmt"
	"os/exec"
	"runtime"
	"sync"
)

// PingTarget checks if an IP is alive
func PingTarget(ip string) bool {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("ping", "-n", "1", "-w", "500", ip)
	} else {
		cmd = exec.Command("ping", "-c", "1", "-W", "1", ip)
	}
	err := cmd.Run()
	return err == nil
}

// DiscoverHosts scans a /24 subnet (e.g., 192.168.1.0-255)
func DiscoverHosts(baseIP string) []string {
	var aliveHosts []string
	var wg sync.WaitGroup
	var mu sync.Mutex

	fmt.Printf("[📡] Discovering active devices on %s.0/24...\n", baseIP)

	for i := 1; i < 255; i++ {
		wg.Add(1)
		targetIP := fmt.Sprintf("%s.%d", baseIP, i)
		go func(ip string) {
			defer wg.Done()
			if PingTarget(ip) {
				mu.Lock()
				aliveHosts = append(aliveHosts, ip)
				mu.Unlock()
				fmt.Printf("    [+] Found Active Device: %s\n", ip)
			}
		}(targetIP)
	}
	wg.Wait()
	return aliveHosts
}
