package utils

import (
	"encoding/json"
	"os"
)

type Config struct {
	Scout struct {
		Target    string `json:"target"`
		StartPort int    `json:"start_port"`
		EndPort   int    `json:"end_port"`
	} `json:"scout"`
	Trading struct {
		Symbol          string `json:"symbol"`
		IntervalSeconds int    `json:"interval_seconds"`
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
	if err != nil {
		return cfg, err
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&cfg)
	return cfg, err
}
