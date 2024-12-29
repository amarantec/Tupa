package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// LoadProjectNameFromGoMod reads the project name (module name) from the go.mod file.
func LoadProjectNameFromGoMod() (string, error) {
	goModPath := "go.mod" // Assuming go.mod is in the project root

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
