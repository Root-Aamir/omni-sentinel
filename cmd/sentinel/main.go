package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/Root-Aamir/omni-sentinel/pkg/scanner"
)

// Task interface allow karta hai kisi bhi naye module ko engine mein fit karna
type Task interface {
	Execute() error
	Name() string
}

func main() {
	fmt.Println("=========================================")
	fmt.Println("🚀 OMNI-SENTINEL v2.0 | Pro Engine Active")
	fmt.Println("Author: Aamir Akram | System: XauCore Engine")
	fmt.Println("Date:", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("=========================================")

	// Tasks list (Aap yahan jitne chahe modules add kar sakte hain)
	tasks := []Task{
		scanner.Scout{
			Target:    "8.8.8.8",
			StartPort: 50,
			EndPort:   100,
		},
		// Future Trading Module yahan aayega
	}

	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go func(t Task) {
			defer wg.Done()
			fmt.Printf("[*] Launching: %s\n", t.Name())
			if err := t.Execute(); err != nil {
				fmt.Printf("[ERROR] %s failed: %v\n", t.Name(), err)
			}
		}(task)
	}

	wg.Wait()
	fmt.Println("\n[✔] All modules executed. Engine standing by.")
}
