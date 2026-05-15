package main

import (
	"fmt"
	"sync"

	"github.com/Root-Aamir/omni-sentinel/pkg/controller"
	"github.com/Root-Aamir/omni-sentinel/pkg/scanner"
	"github.com/Root-Aamir/omni-sentinel/pkg/trading"
	"github.com/Root-Aamir/omni-sentinel/pkg/utils"
)

type Task interface {
	Execute() error
	Name() string
}

func main() {
	cfg, err := utils.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	tasks := []Task{
		scanner.Scout{
			Targets:   cfg.Scout.Targets,
			StartPort: cfg.Scout.StartPort,
			EndPort:   cfg.Scout.EndPort,
			TeleToken: cfg.Telegram.Token,
			TeleID:    cfg.Telegram.ChatID,
		},
		trading.GoldWatcher{
			Symbol:     cfg.Trading.Symbol,
			TeleToken:  cfg.Telegram.Token,
			TeleID:     cfg.Telegram.ChatID,
			TargetBuy:  cfg.Trading.TargetBuy,
			TargetSell: cfg.Trading.TargetSell,
		},
		controller.BotController{
			Token:  cfg.Telegram.Token,
			ChatID: cfg.Telegram.ChatID,
		},
	}

	var wg sync.WaitGroup
	fmt.Println("--- OMNI-SENTINEL v5.0 (Multi-Target) ---")

	for _, task := range tasks {
		wg.Add(1)
		go func(t Task) {
			defer wg.Done()
			fmt.Printf("[+] Starting Module: %s\n", t.Name())
			t.Execute()
		}(task)
	}

	wg.Wait()
}
