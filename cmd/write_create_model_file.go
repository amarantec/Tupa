package cmd

import "os"

func writeCreateModelFile(createModelFilePath string) error {
  /*
    Code to write in create_model.html
  */

  packageContent := []byte(`
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Add new model</title>
    <link href="css/style.css" rel="stylesheet">
  </head>
  <body>
    <h1>Add a new model</h1>
    <form action="/add" method="POST">
      <label for="title">Title:</label>
      <input type="text" id="title" name="title" required>
      <br>
      <label for="body">Body:</label>
      <textarea id="body" name="body" required></textarea>
      <br>
      <button type="submit">Create</button>
    </form>
  </body>
  <!-- Button to return to model list  -->
  <form action="/list" method="GET">
    <button type="submit">List models</button>
  </form>
</html>
`)

  createModelFile, err := os.OpenFile(createModelFilePath, os.O_WRONLY|os.O_APPEND, 0666)
  if err != nil {
    return err
  }
  defer createModelFile.Close()

  // write in create_model.html
  _, err = createModelFile.Write(packageContent)
  if err != nil {
    return err
  }
  return nil
}

