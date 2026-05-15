package main

import (
	"fmt"
	"sync"

	"github.com/Root-Aamir/omni-sentinel/pkg/controller"
	"github.com/Root-Aamir/omni-sentinel/pkg/scanner"
	"github.com/Root-Aamir/omni-sentinel/pkg/trading"
	"github.com/Root-Aamir/omni-sentinel/pkg/utils"
)

// Task interface ensures all modules have a standard execution flow
type Task interface {
	Execute() error
	Name() string
}

func main() {
	// 1. Load System Configuration
	cfg, err := utils.LoadConfig()
	if err != nil {
		fmt.Printf("[!] Fatal Error: Could not load config: %v\n", err)
		return
	}

	fmt.Println("--- OMNI-SENTINEL v6.5 (Final Architecture) ---")
	fmt.Printf("[⚙️] Environment: Hardened | User: %s\n", "Aamir Akram")

	// 2. Network Discovery Phase
	// Note: Verify your base network with 'ipconfig' (e.g., 192.168.1)
	baseNetwork := "192.168.1"
	activeIPs := scanner.DiscoverHosts(baseNetwork)

	// Target Management: Merge Manual Targets with Discovered ones
	var finalTargets []string
	finalTargets = append(finalTargets, cfg.Scout.Targets...)

	if len(activeIPs) > 0 {
		finalTargets = append(finalTargets, activeIPs...)
		fmt.Printf("[📡] Total scan queue updated: %d targets pending.\n", len(finalTargets))

		// Send Discovery notification to Telegram
		msg := fmt.Sprintf("📡 Discovery Complete!\nNew devices found: %d\nTotal targets in queue: %d", len(activeIPs), len(finalTargets))
		utils.SendTelegramAlert(cfg.Telegram.Token, cfg.Telegram.ChatID, msg)
	} else {
		fmt.Println("[⚠️] Local discovery returned 0 results. Proceeding with manual targets.")
	}

	// 3. Initialize Concurrent Modules
	tasks := []Task{
		// Network Module
		scanner.Scout{
			Targets:   finalTargets,
			StartPort: cfg.Scout.StartPort,
			EndPort:   cfg.Scout.EndPort,
			TeleToken: cfg.Telegram.Token,
			TeleID:    cfg.Telegram.ChatID,
		},
		// Trading Intelligence Module
		trading.GoldWatcher{
			Symbol:     cfg.Trading.Symbol,
			TeleToken:  cfg.Telegram.Token,
			TeleID:     cfg.Telegram.ChatID,
			TargetBuy:  cfg.Trading.TargetBuy,
			TargetSell: cfg.Trading.TargetSell,
		},
		// Command & Control (C2) Module
		controller.BotController{
			Token:  cfg.Telegram.Token,
			ChatID: cfg.Telegram.ChatID,
		},
	}

	// 4. Multi-Threaded Execution Engine
	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go func(t Task) {
			defer wg.Done()
			fmt.Printf("[🚀] Launching Module: %s\n", t.Name())
			if err := t.Execute(); err != nil {
				fmt.Printf("[!] Runtime Error in %s: %v\n", t.Name(), err)
			}
		}(task)
	}

	// Engine will keep running until manually interrupted
	wg.Wait()
}
