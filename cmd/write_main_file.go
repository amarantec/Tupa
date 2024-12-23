package cmd

import (
	"fmt"
	"os"
)

func writeMainFile(mainPath, moduleName string) error {
  /*
    Code to write in main.go
  */

  packageContent := fmt.Sprintf(`package main

import ( 
  "fmt"
  "log"
  "net/http"
  "time"
  "%s/internal/db"
  "%s/internal/handler"
  "%s/internal/handler/routes"
)

func main() {
  db.InitDB()
  
  handler.LoadTemplates()
  mux := routes.SetRoutes()
  mux.Handle("/css/",
    http.StripPrefix("/css/",
      http.FileServer(http.Dir("../../web/css"))))

  server := &http.Server{
    Addr:         "127.0.0.1:8080",
    Handler:      mux,
    ReadTimeout:  10 * time.Second,
    WriteTimeout: 10 * time.Second,
  }

  fmt.Println("Server listen on: " + server.Addr)
  log.Fatal(server.ListenAndServe())
}
`, moduleName, moduleName, moduleName)

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
