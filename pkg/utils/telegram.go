package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Send Alert (Existing)
func SendTelegramAlert(token, chatID, message string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	payload := map[string]string{"chat_id": chatID, "text": "🔔 OMNI-SENTINEL ALERT:\n" + message}
	jsonPayload, _ := json.Marshal(payload)
	_, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	return err
}

// Get Latest Messages (New for C2)
type TelegramUpdate struct {
	Result []struct {
		Message struct {
			Text string `json:"text"`
			Chat struct {
				ID int64 `json:"id"`
			} `json:"chat"`
		} `json:"message"`
	} `json:"result"`
}

func GetLatestMessages(token string) (string, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?offset=-1", token)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var update TelegramUpdate
	if err := json.NewDecoder(resp.Body).Decode(&update); err != nil {
		return "", err
	}

	if len(update.Result) > 0 {
		return update.Result[0].Message.Text, nil
	}
	return "", nil
}
