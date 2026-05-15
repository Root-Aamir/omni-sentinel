package main

import (
	"sync"

	"github.com/Root-Aamir/omni-sentinel/pkg/scanner"
	"github.com/Root-Aamir/omni-sentinel/pkg/trading"
	"github.com/Root-Aamir/omni-sentinel/pkg/utils"
)

type Task interface {
	Execute() error
	Name() string
}

func main() {
	cfg, _ := utils.LoadConfig()
	utils.SaveLog("System", "BOOT", "Engine Started")

	if cfg.Telegram.Enabled {
		utils.SendTelegramAlert(cfg.Telegram.Token, cfg.Telegram.ChatID, "🚀 Engine Online")
	}

	tasks := []Task{
		scanner.Scout{
			Target: cfg.Scout.Target, StartPort: cfg.Scout.StartPort, EndPort: cfg.Scout.EndPort,
			TeleToken: cfg.Telegram.Token, TeleID: cfg.Telegram.ChatID,
		},
		trading.GoldWatcher{Symbol: cfg.Trading.Symbol},
	}

	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go func(t Task) {
			defer wg.Done()
			t.Execute()
		}(task)
	}
	wg.Wait()
}
