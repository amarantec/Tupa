package generate

import (
	"fmt"
	"os"
)

func GenerateModelRepositoryFile(modelRepositoryPath, modelName, projectName, dbDrive string) error {
	/*
		Code to write in model_repository.go
	*/

	var packageContent string
	if dbDrive == "postgres" {
		packageContent = fmt.Sprintf(postgresRepositoryTemplate, modelName, projectName, modelName, modelName, modelName, modelName, modelName)
	} else if dbDrive == "mysql" {
		packageContent = fmt.Sprintf(mysqlRepositoryTemplate, modelName, projectName, modelName, modelName, modelName, modelName, modelName)
	} else if dbDrive == "sqlite3" || dbDrive == "" {
		packageContent = fmt.Sprintf(sqlite3RepositoryTemplate, modelName, projectName, modelName, modelName, modelName, modelName, modelName)
	} else {
		return fmt.Errorf("unsupported database driver: %s", dbDrive)
	}

	modelRepositoryFile, err := os.OpenFile(modelRepositoryPath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	_, err = modelRepositoryFile.Write([]byte(packageContent))
	if err != nil {
		return err
	}

	return nil
}
