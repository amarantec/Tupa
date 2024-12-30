package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/amarantec/tupa/constants"
)

// LoadProjectNameFromGoMod reads the project name (module name) from the go.mod file.
func LoadProjectNameFromGoMod() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return constants.EMPTY_STRING, err
	}

	goModPath, err := findGoMod(currentDir)
	if err != nil {
		return constants.EMPTY_STRING, err
	}

	file, err := os.Open(goModPath)
	if err != nil {
		return "", fmt.Errorf("failed to open go.mod: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading go.mod: %w", err)
	}

	return "", fmt.Errorf("module declaration not found in go.mod")
}

func findGoMod(startDir string) (string, error) {
	currentDir := startDir

	for {
		// Caminho completo para o possível arquivo go.mod
		goModPath := filepath.Join(currentDir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			return goModPath, nil
		}

		// Move para o diretório pai
		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			// Se chegarmos na raiz do sistema e não encontramos go.mod
			break
		}
		currentDir = parentDir
	}

	return constants.EMPTY_STRING, fmt.Errorf("go.mod not found starting from %s", startDir)
}
