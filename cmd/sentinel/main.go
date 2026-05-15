package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/Root-Aamir/omni-sentinel/pkg/scanner"
	"github.com/Root-Aamir/omni-sentinel/pkg/trading"
)

// Task interface: Ye hamare engine ka 'Contract' hai.
// Jo bhi naya module (Security ya Trading) banega, usme ye do cheezein honi chahiye.
type Task interface {
	Execute() error
	Name() string
}

func main() {
	// Professional Header
	fmt.Println("=========================================================")
	fmt.Println("🚀 OMNI-SENTINEL v2.5 | Multi-Engine Intelligence")
	fmt.Println("Author: Aamir Akram | Mode: Specialist")
	fmt.Println("Status: Active | Time:", time.Now().Format("15:04:05"))
	fmt.Println("=========================================================")

	// Registry: Yahan hum apne modules ko load karte hain.
	// Aap parallel mein jitne chahe modules chala sakte hain.
	tasks := []Task{
		scanner.Scout{
			Target:    "scanme.nmap.org",
			StartPort: 20,
			EndPort:   85,
		},
		trading.GoldWatcher{
			Symbol: "XAU/USD (Gold)",
		},
	}

	// WaitGroup ensures ke main program tab tak na ruke jab tak saare tasks khatam na ho jayein.
	var wg sync.WaitGroup

	for _, task := range tasks {
		wg.Add(1)

		// Goroutine: Har task ko alag 'Thread' par parallel chala raha hai.
		go func(t Task) {
			defer wg.Done()

			fmt.Printf("[*] Launching: %s\n", t.Name())

			if err := t.Execute(); err != nil {
				fmt.Printf("[!] ERROR in %s: %v\n", t.Name(), err)
			}
		}(task)
	}

	// Engine sabka wait karega
	wg.Wait()

	fmt.Println("=========================================================")
	fmt.Println("✅ SYSTEM STANDBY: All concurrent tasks completed.")
	fmt.Println("=========================================================")
}
