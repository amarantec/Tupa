package generate

import (
	"fmt"
	"os"
)

func WriteMainFile(mainPath, moduleName, dbDrive string) error {
	/*
	   Code to write in main.go
	*/

	var packageContent string
	if dbDrive == "sqlite3" || dbDrive == "" {
		packageContent = fmt.Sprintf(mainFilePackageContentWithSqlite, moduleName, moduleName, moduleName)
	} else if dbDrive == "postgres" {
		packageContent = fmt.Sprintf(mainFilePackageContentWithPostgres, moduleName, moduleName, moduleName)
	} else if dbDrive == "mysql" {
		packageContent = fmt.Sprintf(mainFilePackageContentWithMySql, moduleName, moduleName, moduleName, moduleName)
	} else {
		return fmt.Errorf("unsupported database driver: %s", dbDrive)
	}

	mainFile, err := os.OpenFile(mainPath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer mainFile.Close()

	_, err = mainFile.Write([]byte(packageContent))
	if err != nil {
		return err
	}
	return nil
}
