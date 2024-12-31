package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/amarantec/tupa/cmd/utils"
)

func GenerateModelWebHTML(modelName string, fields []string) error {

	webDir, err := utils.FindWebDirectory()
	if err != nil {
		return err
	}

	htmlCreateFileName := modelName + "_create_template.html"
	modelTemplateWebCreatePath := filepath.Join(webDir, "templates", htmlCreateFileName)
	htmlCreateFile, err := os.Create(modelTemplateWebCreatePath)
	if err != nil {
		return err
	}
	defer htmlCreateFile.Close()

	if err := generateModelWebHTMLCreate(modelTemplateWebCreatePath, modelName, fields); err != nil {
		return err
	}

	htmlListFileName := modelName + "_list_template.html"
	modelTemplateWebListPath := filepath.Join(webDir, "templates", htmlListFileName)
	htmlListFile, err := os.Create(modelTemplateWebListPath)
	if err != nil {
		return err
	}
	defer htmlListFile.Close()

	if err := generateModelWebHTMLList(modelTemplateWebListPath, modelName, fields); err != nil {
		return err
	}

	htmlUpdateFileName := modelName + "_update_template.html"
	modelTemplateWebUpdatePath := filepath.Join(webDir, "templates", htmlUpdateFileName)
	htmlUpdateFile, err := os.Create(modelTemplateWebUpdatePath)
	if err != nil {
		return err
	}
	defer htmlUpdateFile.Close()

	if err := generateModelWebHTMLUpdate(modelTemplateWebUpdatePath, modelName, fields); err != nil {
		return err
	}

	htmlGetFileName := modelName + "_get_template.html"
	modelTemplateWebGetPath := filepath.Join(webDir, "templates", htmlGetFileName)
	htmlGetFile, err := os.Create(modelTemplateWebGetPath)
	if err != nil {
		return err
	}
	defer htmlGetFile.Close()

	if err := generateModelWebHTMLGet(modelTemplateWebGetPath, modelName, fields); err != nil {
		return err
	}
	return nil
}

func generateModelWebHTMLCreate(modelTemplateWebPath string, modelName string, fields []string) error {
	var htmlFields strings.Builder

	for _, field := range fields {
		parts := strings.Split(field, ":")
		if len(parts) < 2 {
			return fmt.Errorf("invalid format of field: %s", field)
		}
		fieldName, fieldType := parts[0], parts[1]
		htmlFields.WriteString(generateHTMLField(fieldName, fieldType) + "\n")
	}

	packageContent := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>Create %s</title>
		<link href="css/style.css" rel="stylesheet">
	</head>
	<body>
		<h1> Create New %s</h1>
		<form action="/create-%s" method="POST">
%s 
			<button type="submit">Create</button>
		</form>
	</body>
</html>
`, modelName, modelName, strings.ToLower(modelName), htmlFields.String())

	modelWebHTMLCreateFile, err := os.OpenFile(modelTemplateWebPath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	_, err = modelWebHTMLCreateFile.Write([]byte(packageContent))
	if err != nil {
		return err
	}

	return nil
}

func generateModelWebHTMLList(modelTemplateWebPath string, modelName string, fields []string) error {
	var htmlFields strings.Builder

	for _, field := range fields {
		parts := strings.Split(field, ":")
		fieldName := parts[0]
		htmlFields.WriteString(fmt.Sprintf("<p>{{ . %s }}</p>\n", fieldName))
	}

	packageContent := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>List All %s</title>
		<link href="css/style.css" rel="stylesheet">
	</head>
	<body>
		<h1>All %s</h1>
		{{ range .Items }}
			%s 
			<hr>
		{{ end }}
	</body>
</html>`, modelName, modelName, htmlFields.String())
	modelWebHTMLListFile, err := os.OpenFile(modelTemplateWebPath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	_, err = modelWebHTMLListFile.Write([]byte(packageContent))
	if err != nil {
		return err
	}

	return nil
}

func generateModelWebHTMLUpdate(modelTemplateWebPath string, modelName string, fields []string) error {
	var htmlFields strings.Builder

	for _, field := range fields {
		parts := strings.Split(field, ":")
		fieldName, fieldType := parts[0], parts[1]
		htmlFields.WriteString(fmt.Sprintf(`<label for="%s">%s</label>
<input type="%s" id="%s" name="%s" value="{{ .%s }}"><br>`, fieldName, fieldName, mapFieldTypeToInputType(fieldType), fieldName, fieldName, fieldName))
	}

	packageContent := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Update %s</title>
	<link href="css/style.css" rel="stylesheet">
  </head>
  <body>
    <h1>Update %s</h1>
    <form action="/update-%s" method="POST">
%s
      <button type="submit">Update</button>
    </form>
  </body>
</html>`, modelName, modelName, strings.ToLower(modelName), htmlFields.String())

	modelWebHTMLUpdateFile, err := os.OpenFile(modelTemplateWebPath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	_, err = modelWebHTMLUpdateFile.Write([]byte(packageContent))
	if err != nil {
		return err
	}

	return nil
}

func generateModelWebHTMLGet(modelTemplateWebPath string, modelName string, fields []string) error {
	var htmlFields strings.Builder

	for _, field := range fields {
		parts := strings.Split(field, ":")
		fieldName := parts[0]
		htmlFields.WriteString(fmt.Sprintf("<p>{{ .%s }}</p>\n", fieldName))
	}
	packageContent := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
    <title>%s Details</title>
	<link href="css/style.css" rel="stylesheet">
  </head>
  <body>
    <h1>%s Details</h1>
%s
  </body>
</html>`, modelName, modelName, htmlFields.String())

	modelWebHTMLGetFile, err := os.OpenFile(modelTemplateWebPath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	_, err = modelWebHTMLGetFile.Write([]byte(packageContent))
	if err != nil {
		return err
	}

	return nil
}

func generateHTMLField(fieldName string, fieldType string) string {
	switch fieldType {
	case "key", "int":
		return fmt.Sprintf(`<label for="%s">%s</label>
<input type="number" id="%s" name="%s"><br>`, fieldName, fieldName, fieldName, fieldName)
	case "string":
		return fmt.Sprintf(`label for="%s">%s</label>
<input type="text" id="%s" name="%s"><br>`, fieldName, fieldName, fieldName, fieldName)
	case "text":
		return fmt.Sprintf(`label for="%s">%s</label>
<textarea id="%s" name="%s"></textarea"><br>`, fieldName, fieldName, fieldName, fieldName)
	default:
		return ""
	}
}

func mapFieldTypeToInputType(fieldType string) string {
	switch fieldType {
	case "key", "int":
		return "number"
	case "string":
		return "text"
	case "text":
		return "textarea"
	default:
		return "text"
	}
}
