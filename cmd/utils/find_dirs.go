package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func FindProjectInternal() (string, error) {
	// Caminha para encontrar o diretório que contém a pasta `internal`.
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory: %w", err)
	}
	for {
		internalPath := filepath.Join(currentDir, "internal")
		if _, err := os.Stat(internalPath); err == nil {
			return internalPath, nil
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			break
		}
		currentDir = parentDir
	}

	return "", fmt.Errorf("internal dir not found")
}

func FindProjectHandler() (string, error) {
	internalDir, err := FindProjectInternal()
	if err != nil {
		return "", fmt.Errorf("failed to get internal directory")
	}

	for {
		handlerPath := filepath.Join(internalDir, "handler")
		if _, err := os.Stat(handlerPath); err == nil {
			return handlerPath, nil
		}

		parentDir := filepath.Dir(internalDir)
		if parentDir == internalDir {
			break
		}
		internalDir = parentDir
	}
	return "", fmt.Errorf("handler dir not found")
}

func FindEnvFile() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory: %w", err)
	}

	for {
		envPath := filepath.Join(currentDir, "config", ".env")
		if _, err := os.Stat(envPath); err == nil {
			return envPath, nil
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			break
		}
		currentDir = parentDir
	}
	return "", fmt.Errorf(".env file not found in config directory")
}
