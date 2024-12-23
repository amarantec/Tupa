package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func createNewProject(projectName, projectPath string) {
  if projectName == "" {
    fmt.Println("You must provide a name for the project.")
    return
  }

  if projectPath == "" {
    fmt.Println("You must provide a path for the project.")
    return
  }

  fmt.Printf("Creating project: %s\n", projectName)

  globalPath := filepath.Join(projectPath, projectName)

  if _, err := os.Stat(globalPath); err == nil {
    fmt.Println("Project directory already exists.")
    return
  }

  /*
      Create project directory by project name
      Create main project directory
  */

  if err := os.Mkdir(globalPath, os.ModePerm); err != nil {
    log.Fatal(err)
  }


  /*
    Start golang project and create go.mod
  */

  startGo := exec.Command("go", "mod", "init", projectName)
  startGo.Dir = globalPath
  startGo.Stdout = os.Stdout
  startGo.Stderr = os.Stderr
  err := startGo.Run()
  if err != nil {
    log.Fatal(err)
  }

  cmdPath := filepath.Join(globalPath, "cmd")
  if err := os.Mkdir(cmdPath, os.ModePerm); err != nil {
    log.Fatal(err)
  }

  webPath := filepath.Join(globalPath, "web")
  if err := os.Mkdir(webPath, os.ModePerm); err != nil {
    log.Fatal(err)
  }

  templatesWeb := filepath.Join(webPath, "templates")
  if err := os.Mkdir(templatesWeb, os.ModePerm); err != nil {
    log.Fatal(err)
  }

  cssWeb := filepath.Join(webPath, "css")
  if err := os.Mkdir(cssWeb, os.ModePerm); err != nil {
    log.Fatal(err)
  }

  apiPath := filepath.Join(cmdPath, "api")
  if err := os.Mkdir(apiPath, os.ModePerm); err != nil {
    log.Fatal(err)
  }

  internalPath := filepath.Join(globalPath, "internal")
  if err := os.Mkdir(internalPath, os.ModePerm); err != nil {
    log.Fatal(err)
  }
  
  handlerPath := filepath.Join(internalPath, "handler")
  if err := os.Mkdir(handlerPath, os.ModePerm); err != nil {
    log.Fatal(err)
  }

  routesPath := filepath.Join(handlerPath, "routes")
  fmt.Println(routesPath)
  if err := os.Mkdir(routesPath, os.ModePerm); err != nil {
    log.Fatal(err)
  }

  databasePath := filepath.Join(internalPath, "db")
  if err := os.Mkdir(databasePath, os.ModePerm); err != nil {
    log.Fatal(err)
  }

  middlewarePath := filepath.Join(internalPath, "middleware")
  if err := os.Mkdir(middlewarePath, os.ModePerm); err != nil {
    log.Fatal(err)
  }

  /*
    Create files
  */

  mainPath := filepath.Join(apiPath, "main.go")
  mainFile, err := os.Create(mainPath)
  if err != nil {
    log.Fatal(err)
  }
  defer mainFile.Close()

  if err := writeMainFile(mainPath, projectName); err != nil {
    log.Fatal(err)
  }

  routesFilePath := filepath.Join(routesPath, "routes.go")
  routesFile, err := os.Create(routesFilePath)
  if err != nil {
    log.Fatal(err)
  }
  defer routesFile.Close()

  if err := writeRoutesFile(routesFilePath, projectName); err != nil {
    log.Fatal(err)
  }

  modelFilePath := filepath.Join(internalPath, "model.go")
  modelFile, err := os.Create(modelFilePath)
  if err != nil {
    log.Fatal(err)
  }
  defer modelFile.Close()

  if err := writeModelFile(modelFilePath, projectName); err != nil {
    log.Fatal(err)
  }

  modelHandlerPath := filepath.Join(handlerPath, "model_handler.go")
  modelHandlerFile, err := os.Create(modelHandlerPath)
  if err != nil {
    log.Fatal(err)
  }
  defer modelHandlerFile.Close()

  if err := writeModelHandlerFile(modelHandlerPath, projectName); err != nil {
    log.Fatal(err)
  }

  dbFilePath := filepath.Join(databasePath, "db.go")
  dbFile, err := os.Create(dbFilePath)
  if err != nil {
    log.Fatal(err)
  }
  defer dbFile.Close()

  if err := writeDbFile(dbFilePath); err != nil {
    log.Fatal(err)
  }

  createModelFilePath := filepath.Join(templatesWeb, "create-model.html")
  createModelFile, err := os.Create(createModelFilePath)
  if err != nil {
    log.Fatal(err)
  }
  defer createModelFile.Close()

  if err := writeCreateModelFile(createModelFilePath); err != nil {
    log.Fatal(err)
  }

  listModelFilePath := filepath.Join(templatesWeb, "list-model.html")
  listModelFile, err := os.Create(listModelFilePath)
  if err != nil {
    log.Fatal(err)
  }
  defer listModelFile.Close()

  if err := writeListModelFile(listModelFilePath); err != nil {
    log.Fatal(err)
  }

  templatesFilePath := filepath.Join(handlerPath, "templates.go")
  templatesFile, err := os.Create(templatesFilePath)
  if err != nil {
    log.Fatal(err)
  }
  defer templatesFile.Close()

  if err := writeTemplatesFile(templatesFilePath); err != nil {
    log.Fatal(err)
  }
}
