package generate

import (
	"fmt"
	"os"
)

func GenerateModelServiceFile(modelServicePath, modelName, projectName string) error {
	/*
		Code to write in model_service.go
	*/
	packageContent := fmt.Sprintf(`package %s

import (
	"context"
	"%s/internal"
)

type I%sService interface{}

type %sService struct {
	%sRepositoryService %sRepository
}

func New%sService(repo %sRepository) I%sService {
	return &%sService{%sRepositoryService: repo}
}

`, modelName, projectName, modelName, modelName, modelName, modelName,
		modelName, modelName, modelName, modelName, modelName)

	modelServiceFile, err := os.OpenFile(modelServicePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	_, err = modelServiceFile.Write([]byte(packageContent))
	if err != nil {
		return err
	}
	return nil
}
