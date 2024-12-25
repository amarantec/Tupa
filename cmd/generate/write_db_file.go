package generate

import (
	"fmt"
	"os"
)

func WriteDbFile(dbFilePath, moduleName string) error {
	/*
	   Code to write in db.go
	*/

	packageContent := fmt.Sprintf(`package db

import (
  "database/sql"
  "fmt"
  _"github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
 var err error
 DB, err = sql.Open("sqlite3", "api.db")
  if err != nil {
    panic("Could not connect to dabase.")
  }

  DB.SetMaxOpenConns(10)
  DB.SetMaxIdleConns(5)
  createTables()
}

/*
  Function to create a "model" table for example when running the project for the first time.
*/
func createTables() {
  create%sTable := " CREATE TABLE IF NOT EXISTS %ss (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, body TEXT NOT NULL);"
  _, err := DB.Exec(create%sTable)
  if err != nil {
    fmt.Println(err)
    panic("Could not create %ss table.")
  }
}
`, moduleName, moduleName, moduleName, moduleName)

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
