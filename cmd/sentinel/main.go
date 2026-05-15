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

	tasks := []Task{
		scanner.Scout{
			Target: cfg.Scout.Target, StartPort: cfg.Scout.StartPort, EndPort: cfg.Scout.EndPort,
			TeleToken: cfg.Telegram.Token, TeleID: cfg.Telegram.ChatID,
		},
		trading.GoldWatcher{
			Symbol:    cfg.Trading.Symbol,
			TeleToken: cfg.Telegram.Token, TeleID: cfg.Telegram.ChatID,
			TargetBuy: cfg.Trading.TargetBuy, TargetSell: cfg.Trading.TargetSell,
		},
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
