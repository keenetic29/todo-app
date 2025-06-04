package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	DBURL         string
	ServerAddress string
	JWTSecret     string
	LogFile       string
}

func LoadConfig(filename string) (*Config, error) {
	if err := loadEnvFile(filename); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	cfg := &Config{
		DBURL:         getEnv("DB_URL", ""),
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"), 
		JWTSecret:     getEnv("JWT_SECRET", ""),
		LogFile:       getEnv("LOG_FILE", "app.log"),
	}

	if cfg.DBURL == "" {
		return nil, fmt.Errorf("DB_URL is required")
	}
	if cfg.JWTSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET is required")
	}

	return cfg, nil
}

func loadEnvFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err 
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue 
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue 
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		if key != "" {
			os.Setenv(key, value)
		}
	}

	return scanner.Err()
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}