package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetDBDriverFromEnv() (string, error) {
	envPath, err := FindEnvFile()
	if err != nil {
		return "", err
	}

	file, err := os.Open(envPath)
	if err != nil {
		return "", fmt.Errorf("failed to open .env file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "DB_DRIVER=") {
			return strings.TrimSpace(strings.TrimPrefix(line, "DB_DRIVER=")), nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("failed to read .env file: %w", err)
	}

	return "", fmt.Errorf("DB_DRIVER not found in .env file")
}
