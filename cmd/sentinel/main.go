package main

import (
	"fmt"
	"sync"

	"github.com/Root-Aamir/omni-sentinel/pkg/scanner"
	"github.com/Root-Aamir/omni-sentinel/pkg/trading"
	"github.com/Root-Aamir/omni-sentinel/pkg/utils" // Import utils for logs & config
)

type Task interface {
	Execute() error
	Name() string
}

func main() {
	// 1. Config Load Karein
	cfg, err := utils.LoadConfig()
	if err != nil {
		fmt.Println("❌ Critical Error: Could not load config.json")
		return
	}

	// 2. Boot Log Save Karein (Ab ye logs folder mein file banayega)
	utils.SaveLog("System", "BOOT", "Engine v2.6 started with config.json")

	fmt.Println("=========================================")
	fmt.Println("🚀 OMNI-SENTINEL: CONFIG-DRIVEN ENGINE")
	fmt.Println("=========================================")

	// 3. Config se data utha kar modules initialize karein
	tasks := []Task{
		scanner.Scout{
			Target:    cfg.Scout.Target,
			StartPort: cfg.Scout.StartPort,
			EndPort:   cfg.Scout.EndPort,
		},
		trading.GoldWatcher{
			Symbol: cfg.Trading.Symbol,
		},
	}

	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go func(t Task) {
			defer wg.Done()
			fmt.Printf("[*] Launching: %s\n", t.Name())
			if err := t.Execute(); err != nil {
				utils.SaveLog(t.Name(), "ERROR", err.Error())
			}
		}(task)
	}

	wg.Wait()
	utils.SaveLog("System", "SHUTDOWN", "All tasks completed.")
	fmt.Println("\n[✔] Execution finished. Check logs folder.")
}
