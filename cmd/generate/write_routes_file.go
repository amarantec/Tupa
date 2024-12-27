package generate

import (
	"fmt"
	"os"
)

func WriteRoutesFile(routesFilePath, moduleName string) error {
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

  m.HandleFunc("/", handler.LoadIndexTemplate)

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
