package utils

import (
	"encoding/json"
	"os"
)

type Config struct {
	Scout struct {
		Targets   []string `json:"targets"` // Badla hua: Ab ye list hai
		StartPort int      `json:"start_port"`
		EndPort   int      `json:"end_port"`
	} `json:"scout"`
	Trading struct {
		Symbol          string  `json:"symbol"`
		IntervalSeconds int     `json:"interval_seconds"`
		TargetBuy       float64 `json:"target_buy"`
		TargetSell      float64 `json:"target_sell"`
	} `json:"trading"`
	Telegram struct {
		Token   string `json:"token"`
		ChatID  string `json:"chat_id"`
		Enabled bool   `json:"enabled"`
	} `json:"telegram"`
}

func LoadConfig() (Config, error) {
	var cfg Config
	file, err := os.Open("config.json")
	if err == nil {
		defer file.Close()
		json.NewDecoder(file).Decode(&cfg)
	}

	// Environment variables override
	envToken := os.Getenv("TELE_TOKEN")
	envID := os.Getenv("TELE_ID")

	if envToken != "" {
		cfg.Telegram.Token = envToken
		cfg.Telegram.Enabled = true
	}
	if envID != "" {
		cfg.Telegram.ChatID = envID
	}

	return cfg, nil
}
