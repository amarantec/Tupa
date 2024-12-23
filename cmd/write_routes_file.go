package cmd

import (
	"fmt"
	"os"
)

func writeRoutesFile(routesFilePath, moduleName string) error {
  /*
    Code to write in routes.go
  */

  packageContent := fmt.Sprintf(`package routes
  
import(
  "net/http"
  "%s/internal/handler"
)

func SetRoutes() *http.ServeMux{
  m := http.NewServeMux()

  m.HandleFunc("/create-new-model", handler.LoadCreateModelForm)
  m.HandleFunc("/add", handler.CreateNewModel)
  m.HandleFunc("/get-all-models", handler.GetAllModels)
  m.HandleFunc("/list", handler.ListModels)

  return m
}
`, moduleName)
   routesFile, err := os.OpenFile(routesFilePath, os.O_WRONLY|os.O_APPEND, 0666)
  if err != nil {
    return err
  }
  defer routesFile.Close()

  // write in routes.go
  _, err = routesFile.Write([]byte(packageContent))
  if err != nil {
    return err
  }
  return nil
}
