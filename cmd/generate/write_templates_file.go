package generate

import "os"

func WriteTemplatesFile(templatesFilePath string) error {
	/*
	   Code to write in templates.go
	*/

	packageContent := []byte(`package handler
import (
  "log"
  "text/template"
)

var Templates *template.Template

func LoadTemplates() {
  var err error
  Templates, err = template.ParseGlob("./web/templates/*html")
  if err != nil {
    log.Fatalf("Error loading templates: %v\n", err)
  }
}
`)

	templatesFile, err := os.OpenFile(templatesFilePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer templatesFile.Close()

	// write in templates.go

	_, err = templatesFile.Write(packageContent)
	if err != nil {
		return err
	}
	return nil
}
