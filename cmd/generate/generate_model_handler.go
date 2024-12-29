package generate

import (
	"fmt"
	"os"
)

func GenerateModelHandlerFile(modelHandlerPath, modelName, projectName string) error {
	/*
		Code to write in model_handler.go
	*/

	packageContent := fmt.Sprintf(`package %sHandler

import (
	"%s/internal/%s"
)

type %sHandler struct {
	svc %s.I%sService
}

func New%sHandler(service %s.I%sService) *%sHandler {
	return &%sHandler{svc: service}
}
`, modelName, projectName, modelName, modelName, modelName, modelName,
		modelName, modelName, modelName, modelName, modelName,
	)

	modelHandlerFile, err := os.OpenFile(modelHandlerPath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	_, err = modelHandlerFile.Write([]byte(packageContent))
	if err != nil {
		return err
	}
	return nil
}
