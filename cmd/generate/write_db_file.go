package generate

import (
	"fmt"
	"os"

	"github.com/amarantec/tupa/constants"
)

func WriteDbFile(dbFilePath, moduleName, dbDrive string) error {
	/*
	   Code to write in db.go
	*/

	var packageContent string

	if dbDrive == "postgres" {
		packageContent = postgresDatabasePackageContent
	} else if dbDrive == "mysql" {
		packageContent = mysqlDatabasePackageContent
	} else if dbDrive == constants.EMPTY_STRING || dbDrive == "sqlite3" {
		packageContent = sqlite3DatabasePackageContent
	} else {
		fmt.Println("Unsupported database driver: " + dbDrive)
		os.Exit(1)
	}

	dbFile, err := os.OpenFile(dbFilePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer dbFile.Close()

	// write in db.go
	_, err = dbFile.Write([]byte(packageContent))
	if err != nil {
		return err
	}
	return nil
}
