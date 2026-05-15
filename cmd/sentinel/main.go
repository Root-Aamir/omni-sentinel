package main

import (
	"fmt"
	"time"

	"github.com/Root-Aamir/omni-sentinel/pkg/scanner" // Naya module import kiya
)

func main() {
	fmt.Println("-----------------------------------------")
	fmt.Println("🚀 OMNI-SENTINEL: Core Engine Started")
	fmt.Println("User: Aamir Akram | System: Pro-Specialist")
	fmt.Println("Time:", time.Now().Format("15:04:05"))
	fmt.Println("-----------------------------------------")

	fmt.Println("[+] Running Scout Module: Scanning Target...")

	// Example: Google ke common ports check karte hain
	target := "8.8.8.8"
	ports := []int{53, 80, 443}

	for _, port := range ports {
		scanner.ScanPort(target, port)
	}

	fmt.Println("[+] Scan Completed.")
}
