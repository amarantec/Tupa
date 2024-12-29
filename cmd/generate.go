package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/amarantec/tupa/cmd/generate"
	"github.com/amarantec/tupa/cmd/utils"
)

func generateNewModel(args []string) error {
	if len(args) < 2 {
		fmt.Println(args)
		return fmt.Errorf("usage: tupa generate -n <ModelName> <Field:Type> <Field:Type>")
	}

	if err := modelNewStruct(args); err != nil {
		return err
	}

	modelName := args[0]
	internalPath, err := utils.FindProjectInternal()
	if err != nil {
		return err
	}

	projectName, err = utils.LoadProjectNameFromGoMod()
	if err != nil {
		return fmt.Errorf("failed to get module name from go.mod")
	}

	dbDrive, err = utils.GetDBDriverFromEnv()
	if err != nil {
		return err
	}

	modelPath := filepath.Join(internalPath, modelName)
	if err := os.Mkdir(modelPath, os.ModePerm); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	modelServiceFileName := strings.ToLower(modelName + "_service.go")
	modelServicePath := filepath.Join(modelPath, modelServiceFileName)
	modelServiceFile, err := os.Create(modelServicePath)
	if err != nil {
		log.Fatal(err)
	}
	defer modelServiceFile.Close()

	if err := generate.GenerateModelServiceFile(modelServicePath, modelName, projectName); err != nil {
		log.Fatal(err)
	}

	modelRepositoryFileName := strings.ToLower(modelName + "_repository.go")
	modelRepositoryPath := filepath.Join(modelPath, modelRepositoryFileName)
	modelRepositoryFile, err := os.Create(modelRepositoryPath)
	if err != nil {
		log.Fatal(err)
	}
	defer modelRepositoryFile.Close()

	if err := generate.GenerateModelRepositoryFile(modelRepositoryPath, modelName, projectName, dbDrive); err != nil {
		log.Fatal(err)
	}

	handlerPath, err := utils.FindProjectHandler()
	if err != nil {
		return err
	}

	modelHandler := filepath.Join(handlerPath, modelName)
	modelHandlerDirName := modelHandler + "Handler"
	if err := os.Mkdir(modelHandlerDirName, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	modelHandlerFileName := strings.ToLower(modelName + "_handler.go")
	modelHandlerPath := filepath.Join(modelHandlerDirName, modelHandlerFileName)
	modelHandlerFile, err := os.Create(modelHandlerPath)
	if err != nil {
		log.Fatal(err)
	}
	defer modelHandlerFile.Close()

	if err := generate.GenerateModelHandlerFile(modelHandlerPath, modelName, projectName); err != nil {
		log.Fatal(err)
	}

	return nil
}
