package utils

import (
	"encoding/json"
	"os"
)

// 1. Struct Definition (Pehle ye hona zaroori hai)
type Config struct {
	Scout struct {
		Target    string `json:"target"`
		StartPort int    `json:"start_port"`
		EndPort   int    `json:"end_port"`
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

// 2. Function Definition
func LoadConfig() (Config, error) {
	var cfg Config

	// JSON file read karein
	file, err := os.Open("config.json")
	if err == nil {
		defer file.Close()
		json.NewDecoder(file).Decode(&cfg)
	}

	// System Environment se secrets uthayein (Security)
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
