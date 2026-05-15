package controller

import (
	"fmt"
	"strings"
	"time"

	"github.com/Root-Aamir/omni-sentinel/pkg/utils"
)

type BotController struct {
	Token  string
	ChatID string
}

func (b BotController) Name() string { return "Command-Center" }

func (b BotController) Execute() error {
	fmt.Println("[🤖] Command Center Active. Listening for instructions...")

	// lastMsgID use karte hain taaki purane messages repeat na hon
	lastMsg := ""

	for {
		// Telegram se naya message fetch karna
		msg, err := utils.GetLatestMessages(b.Token)

		// Agar message naya hai aur empty nahi hai
		if err == nil && msg != "" && msg != lastMsg {
			lastMsg = msg
			fmt.Printf("[Telegram] Received Command: %s\n", msg)

			// Command Processing Logic
			switch {
			case strings.HasPrefix(msg, "/status"):
				statusReport := "🛡️ **SENTINEL V6.0 STATUS**\n" +
					"--------------------------\n" +
					"🟢 Engine: Running\n" +
					"📡 Discovery: Active\n" +
					"📈 Trading: XAU/USD Monitoring\n" +
					"🔒 Mode: Hardened (Env-Only)"
				utils.SendTelegramAlert(b.Token, b.ChatID, statusReport)

			case strings.HasPrefix(msg, "/price"):
				priceUpdate := "📊 **MARKET STRATEGY**\n" +
					"Asset: XAU/USD (Gold)\n" +
					"Target BUY: < $2340\n" +
					"Target SELL: > $2360\n" +
					"Status: Waiting for Entry..."
				utils.SendTelegramAlert(b.Token, b.ChatID, priceUpdate)

			case strings.HasPrefix(msg, "/discover"):
				utils.SendTelegramAlert(b.Token, b.ChatID, "🔎 **NETWORK SCAN TRIGGERED**\nStarting subnet discovery on local Wi-Fi. Check Terminal for live host logs.")
				// Future: Hum yahan channel ke through scanner module ko trigger karenge

			case strings.HasPrefix(msg, "/help"):
				menu := "🕹️ **SENTINEL C2 MENU**\n" +
					"/status - System health check\n" +
					"/price - Current trading targets\n" +
					"/discover - Manually scan network\n" +
					"/help - Show this control panel"
				utils.SendTelegramAlert(b.Token, b.ChatID, menu)

			default:
				// Agar koi unknown command aaye
				if strings.HasPrefix(msg, "/") {
					utils.SendTelegramAlert(b.Token, b.ChatID, "⚠️ Unknown Command. Type /help for options.")
				}
			}
		}

		// API rate limiting (3 seconds delay)
		time.Sleep(3 * time.Second)
	}
}
