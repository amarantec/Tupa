package cmd

import "os"

func writeDbFile(dbFilePath string) error {
  /*
    Code to write in db.go
  */

  packageContent := []byte(`package db

import (
  "database/sql"
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
func createTables() {
  createAccountTable := " CREATE TABLE IF NOT EXISTS models (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, body TEXT NOT NULL);"
  _, err := DB.Exec(createAccountTable)
  if err != nil {
    panic("Could not create accounts table.")
  }
}
`)

  dbFile, err := os.OpenFile(dbFilePath, os.O_WRONLY|os.O_APPEND, 0666)
  if err != nil {
    return err
  }
  defer dbFile.Close()

  // write in db.go
  _, err = dbFile.Write(packageContent)
  if err != nil {
    return err
  }
  return nil
}
