package generate

import (
	"fmt"
	"os"
)

func WriteListModelTemplateFile(listModelTemplateFilePath, projectName string) error {
	/*
	   Code to write in list_model.html
	*/

	packageContent := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>List</title>
    <link href="css/style.css" rel="stylesheet">
  </head>
  <body>
    <h1>List all</h1>
    {{ range .%s }}
      <div>
        <h2>{{ .Title }}</h2>
        <p>{{ .Body }}</p>
      </div>
      <hr>
    {{ end }}
    <form action="/create-new-model" method="GET">
        <button type="submit">Add new</button>
    </form> 
  </body>
</html>
`, projectName)

	listModelTemplateFile, err := os.OpenFile(listModelTemplateFilePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer listModelTemplateFile.Close()

	// write in list_model.html
	_, err = listModelTemplateFile.Write([]byte(packageContent))
	if err != nil {
		return err
	}
	return nil
}
