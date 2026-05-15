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
}

func LoadConfig() (Config, error) {
	var cfg Config
	file, err := os.Open("config.json") // Extension .json honi chahiye
	if err != nil {
		return cfg, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&cfg)
	return cfg, err
}
