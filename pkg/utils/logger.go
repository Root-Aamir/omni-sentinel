package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

type LogEntry struct {
	Timestamp string      `json:"timestamp"`
	Module    string      `json:"module"`
	Level     string      `json:"level"`
	Data      interface{} `json:"data"`
}

var mu sync.Mutex // Mutex ensure karta hai ki do modules ek saath file mein na likhein

func SaveLog(module, level string, data interface{}) {
	mu.Lock()
	defer mu.Unlock()

	entry := LogEntry{
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Module:    module,
		Level:     level,
		Data:      data,
	}

	// JSON mein convert karein
	jsonData, _ := json.MarshalIndent(entry, "", "  ")

	// File mein append karein (Daily log file)
	fileName := fmt.Sprintf("logs/%s.json", time.Now().Format("2006-01-02"))
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Logger Error:", err)
		return
	}
	defer f.Close()

	f.WriteString(string(jsonData) + ",\n")
}
