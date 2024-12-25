package generate

import (
	"fmt"

	"github.com/amarantec/tupa/cmd/model"
)

func GenerateNewModel(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: tupa generate <ModelName> <Field:Type> <Field:Type>")
	}

	err := model.ModelNewStruct(args)
	if err != nil {
		return err
	}

	modelName := args[0]
	err = generateFiles(modelName)
	if err != nil {
		return err
	}

	return nil
}
