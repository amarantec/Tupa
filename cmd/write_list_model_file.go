package cmd

import "os"

func writeListModelFile(listModelFilePath string) error {
  /*
    Code to write in list_model.html
  */

  packageContent := []byte(`
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>List models</title>
    <link href="css/style.css" rel="stylesheet">
  </head>
  <body>
    <h1>List all models</h1>
    {{ range .Model }}
      <div>
        <h2>{{ .Title }}</h2>
        <p>{{ .Body }}</p>
      </div>
      <hr>
    {{ end }}
   <!-- Button to add a new model -->
    <form action="/create-new-model" method="GET">
        <button type="submit">Add new model</button>
    </form> 
  </body>
</html>
`)

  listModelFile, err := os.OpenFile(listModelFilePath, os.O_WRONLY|os.O_APPEND, 0666)
  if err != nil {
    return err
  }
  defer listModelFile.Close()

  // write in list_model.html
  _, err = listModelFile.Write(packageContent)
  if err != nil {
    return err
  }
  return nil
}

