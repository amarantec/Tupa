package model

import (
	"fmt"
	"os"
	"strings"
)

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
			return fmt.Errorf("invalid field formart: %s, expected <FieldName>:<Type>", field)
		}

		fieldName := parts[0]
		fieldType := parts[1]
		fieldDefinitions = append(fieldDefinitions, fmt.Sprintf("%s    %s", fieldName, strings.ToLower(fieldType)))
	}

	modelContent := fmt.Sprintf("package internal\n\ntype %s struct {\n  %s\n}\n",
		modelName, strings.Join(fieldDefinitions, "\n  "))

	filePath := fmt.Sprintf("internal/%s.go", strings.ToLower(modelName))

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(modelContent)
	if err != nil {
		return err
	}

	fmt.Printf("Model %s created successfully at: %s\n", modelName, filePath)
	return nil
}
