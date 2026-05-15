package scanner

import (
	"fmt"
	"net"
	"time"
)

// ScanPort function ek specific IP aur Port ko check karti hai
func ScanPort(target string, port int) {
	address := fmt.Sprintf("%s:%d", target, port)
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)

	if err != nil {
		// Port band hai ya filter ho raha hai
		return
	}
	conn.Close()
	fmt.Printf("[!] ALERT: Port %d is OPEN on %s\n", port, target)
}
