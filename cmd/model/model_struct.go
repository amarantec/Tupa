package model

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func findProjectRoot() (string, error) {
	// Caminha para encontrar o diretório que contém a pasta `internal`.
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory: %w", err)
	}

	for {
		if _, err := os.Stat(filepath.Join(currentDir, "internal")); err == nil {
			return currentDir, nil
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			break
		}
		currentDir = parentDir
	}

	return "", fmt.Errorf("project root not found")
}

func ModelNewStruct(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: tupa new model <ModelName> <Field:Type> <Field:Type>")
	}

	modelName := args[0]
	fields := args[1:]

	var fieldDefinitions []string
	for _, field := range fields {
		parts := strings.Split(field, ":")
		if len(parts) != 2 {
			return fmt.Errorf("invalid field format: %s, expected <FieldName>:<Type>", field)
		}

		fieldName := parts[0]
		fieldType := parts[1]
		fieldDefinitions = append(fieldDefinitions, fmt.Sprintf("%s    %s", fieldName, strings.ToLower(fieldType)))
	}

	modelContent := fmt.Sprintf("package internal\n\ntype %s struct {\n  %s\n}\n",
		modelName, strings.Join(fieldDefinitions, "\n  "))

	// Localiza o diretório raiz do projeto.
	projectRoot, err := findProjectRoot()
	if err != nil {
		return err
	}

	// Constrói o caminho para o arquivo.
	modelFilePath := filepath.Join(projectRoot, "internal", fmt.Sprintf("%s.go", strings.ToLower(modelName)))

	// Cria o arquivo de modelo.
	file, err := os.OpenFile(modelFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(modelContent)
	if err != nil {
		return err
	}
	return nil
}
