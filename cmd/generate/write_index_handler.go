package generate

import "os"

func WriteIndexHandlerFile(indexHandlerFilePath string) error {
	/*
		Code to wirte in index_handler.go
	*/

	packageContent := []byte(`package handler
	
import "net/http"
	
func LoadIndexTemplate(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusBadRequest)
		return
	}
}`)

	indexHandlerFile, err := os.OpenFile(indexHandlerFilePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	_, err = indexHandlerFile.Write(packageContent)
	if err != nil {
		return err
	}

	return nil
}
