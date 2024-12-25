package generate

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func generateFiles(modelName string) error {
	file := strings.ToLower(modelName) + ".go"
	modelHandlerPath := filepath.Join("internal/handler/", file)
	modelHandlerFile, err := os.Create(modelHandlerPath)
	if err != nil {
		log.Fatal(err)
	}
	defer modelHandlerFile.Close()

	/*
		if err := writeModelHandler(modelName); err != nil {
			log.Fatal(err)
		}

	*/
	return nil
}
