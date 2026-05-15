package utils

import (
	"fmt"
	"net/http"
	"net/url"
)

func SendTelegramAlert(token, chatID, message string) {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	resp, err := http.PostForm(apiURL, url.Values{
		"chat_id": {chatID},
		"text":    {"🔔 OMNI-SENTINEL ALERT:\n" + message},
	})
	if err != nil {
		fmt.Println("❌ Telegram Error:", err)
		return
	}
	defer resp.Body.Close()
}
