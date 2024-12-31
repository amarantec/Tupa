package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/amarantec/tupa/cmd/generate"
	"github.com/amarantec/tupa/cmd/utils"
	"github.com/amarantec/tupa/constants"
)

// Constantes para valores padrao

// Funcao para gerar um novo modelo
func generateNewModel(args []string) error {
	if len(args) < 2 {
		fmt.Println(args)
		return fmt.Errorf(constants.USAGE_MESSAGE)
	}

	// Gera a estrutura do modelo
	if err := modelNewStruct(args); err != nil {
		return err
	}

	// Variaveis de configuracao
	modelName := args[0]
	fields := args[1:]

	if err := GenerateModelWebHTML(modelName, fields); err != nil {
		return err
	}

	internalPath, err := utils.FindProjectInternal()
	if err != nil {
		return err
	}

	projectName, err := utils.LoadProjectNameFromGoMod()
	if err != nil {
		return fmt.Errorf("failed to get module name from go.mod")
	}

	dbDriver, err := utils.GetDBDriverFromEnv()
	if err != nil {
		return err
	}

	// Caminho do modelo e cricao do diretorio
	modelPath := filepath.Join(internalPath, modelName)
	if err := os.Mkdir(modelPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// Gerar arquivos do modelo
	if err := generateModelFiles(modelName, modelPath, projectName, dbDriver); err != nil {
		log.Fatal(err)
	}

	// Gerar SQL para migracao
	sql, err := GenerateSQLFromStruct(args, dbDriver)
	if err != nil {
		return fmt.Errorf("failed to generate SQL from struct: %w", err)
	}

	// Criar migracao
	if err := CreateMigrations(modelName, sql); err != nil {
		return fmt.Errorf("failed to create migration: %w", err)
	}

	return nil
}

// Função auxiliar para gerar os arquivos do modelo
func generateModelFiles(modelName, modelPath, projectName, dbDriver string) error {
	// Gerar arquivo de servico
	modelServiceFileName := strings.ToLower(modelName + constants.SERVICE_SUFFIX)
	modelServicePath := filepath.Join(modelPath, modelServiceFileName)
	modelServiceFile, err := os.Create(modelServicePath)
	if err != nil {
		return err
	}
	defer modelServiceFile.Close()

	if err := generate.GenerateModelServiceFile(modelServicePath, modelName, projectName); err != nil {
		return err
	}

	// Gerar arquivo de repositorio
	modelRepositoryFileName := strings.ToLower(modelName + constants.REPOSITORY_SUFFIX)
	modelRepositoryPath := filepath.Join(modelPath, modelRepositoryFileName)
	modelRepositoryFile, err := os.Create(modelRepositoryPath)
	if err != nil {
		return err
	}
	defer modelRepositoryFile.Close()

	if err := generate.GenerateModelRepositoryFile(modelRepositoryPath, modelName, projectName, dbDriver); err != nil {
		return err
	}

	// Gerar arquivo de handler
	handlerPath, err := utils.FindProjectHandler()
	if err != nil {
		return err
	}

	modelHandler := filepath.Join(handlerPath, modelName)
	modelHandlerDirName := modelHandler + "Handler"
	if err := os.Mkdir(modelHandlerDirName, os.ModePerm); err != nil {
		return err
	}

	modelHandlerFileName := strings.ToLower(modelName + constants.HANDLER_SUFFIX)
	modelHandlerPath := filepath.Join(modelHandlerDirName, modelHandlerFileName)
	modelHandlerFile, err := os.Create(modelHandlerPath)
	if err != nil {
		return err
	}
	defer modelHandlerFile.Close()

	if err := generate.GenerateModelHandlerFile(modelHandlerPath, modelName, projectName); err != nil {
		return err
	}

	return nil
}
