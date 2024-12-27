package generate

import (
	"fmt"
	"os"
)

func WriteEnvFile(envFilePath, dbDrive string) error {
	var config string

	if dbDrive == "postgres" {
		config = `
		DB_HOST=localhost
		DB_PORT=5432
		POSTGRES_USER=postgres_user
		POSTGRES_PASSWORD=postgres_password
		POSTGRES_DB=app_development
		`
	} else if dbDrive == "mysql" {
		config = `
		DB_HOST=localhost
		DB_PORT=3306
		MYSQL_USER=mysql_user
		MYSQL_PASSWORD=mysql_password
		MYSQL_DATABASE=app_development
		`
	} else if dbDrive == "" || dbDrive == "sqlite3" {
		config = ""
	} else {
		fmt.Println("Unsupported database driver: " + dbDrive)
		os.Exit(1)
	}

	envFile, err := os.OpenFile(envFilePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer envFile.Close()

	_, err = envFile.Write([]byte(config))
	if err != nil {
		return err
	}
	return nil
}
