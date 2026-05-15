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
	fmt.Println("[🤖] Command Center Active. Listening for commands...")
	lastMsg := ""

	for {
		msg, err := utils.GetLatestMessages(b.Token)
		if err == nil && msg != "" && msg != lastMsg {
			lastMsg = msg
			fmt.Printf("[Telegram] Received: %s\n", msg)

			switch {
			case strings.HasPrefix(msg, "/status"):
				utils.SendTelegramAlert(b.Token, b.ChatID, "✅ OMNI-SENTINEL is running.\nLocation: Roorkee Labs\nStatus: Secure")

			case strings.HasPrefix(msg, "/price"):
				utils.SendTelegramAlert(b.Token, b.ChatID, "📊 Current monitoring range:\nBUY: < $2340\nSELL: > $2360")

			case strings.HasPrefix(msg, "/help"):
				menu := "🛠 Sentinel Control Menu:\n/status - Check engine health\n/price - Get current targets\n/help - Show this menu"
				utils.SendTelegramAlert(b.Token, b.ChatID, menu)
			}
		}
		time.Sleep(3 * time.Second) // API Rate limiting se bachne ke liye
	}
}
