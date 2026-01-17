package main

import (
	"bufio"
	"os"
	"strings"
)

// Config holds the color configuration for the TUI
type Config struct {
	WindowColor   string
	UserColor     string
	DateTimeColor string
	MsgColor      string
	TextColor     string
}

// DefaultConfig returns the default configuration
func DefaultConfig() Config {
	return Config{
		WindowColor:   "#ffffff",
		UserColor:     "#AAAAAA",
		DateTimeColor: "#555555",
		MsgColor:      "#ffffff",
		TextColor:     "#ffffff",
	}
}

// LoadConfig reads the configuration from a file
func LoadConfig(path string) (Config, error) {
	config := DefaultConfig()

	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return config, nil // Return default if file doesn't exist
		}
		return config, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "WINDOW":
			config.WindowColor = value
		case "USER":
			config.UserColor = value
		case "DATETIME":
			config.DateTimeColor = value
		case "MSG":
			config.MsgColor = value
		case "TEXT":
			config.TextColor = value
		}
	}

	if err := scanner.Err(); err != nil {
		return config, err
	}

	return config, nil
}
