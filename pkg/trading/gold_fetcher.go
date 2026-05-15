package trading

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Root-Aamir/omni-sentinel/pkg/utils"
)

type GoldWatcher struct {
	Symbol     string
	TeleToken  string
	TeleID     string
	TargetBuy  float64
	TargetSell float64
}

func (g GoldWatcher) Name() string { return "Quant-Intelligence" }

func (g GoldWatcher) Execute() error {
	fmt.Printf("[📈] %s: Monitoring market for targets...\n", g.Symbol)

	// Simulation for 5 ticks
	for i := 0; i < 5; i++ {
		price := 2335.0 + rand.Float64()*(2365.0-2335.0)
		fmt.Printf("    [Market] %s | Price: $%.2f\n", g.Symbol, price)

		// Smart Alert Logic
		if price <= g.TargetBuy {
			utils.SendTelegramAlert(g.TeleToken, g.TeleID, fmt.Sprintf("📉 BUY SIGNAL: %s dropped to $%.2f", g.Symbol, price))
		} else if price >= g.TargetSell {
			utils.SendTelegramAlert(g.TeleToken, g.TeleID, fmt.Sprintf("📈 SELL SIGNAL: %s rose to $%.2f", g.Symbol, price))
		}

		time.Sleep(5 * time.Second)
	}
	return nil
}
