package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/amarantec/tupa/cmd/utils"
)

func modelNewStruct(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: tupa model <ModelName> <Field:Type> <Field:Type>")
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

	projectInternal, err := utils.FindProjectInternal()
	if err != nil {
		return err
	}

	// Constrói o caminho para o arquivo.
	modelFilePath := filepath.Join(projectInternal, fmt.Sprintf("%s.go", strings.ToLower(modelName)))

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
