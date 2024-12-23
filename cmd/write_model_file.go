package cmd

import (
	"fmt"
	"os"
)

func writeModelFile(modelFilePath, moduleName string) error {
  /*
    Code to write in model.go
  */

  packageContent := fmt.Sprintf(`package internal

import (
  "context"
  "%s/internal/db"
)

type Model struct {
  Id      int64
  Title   string   
  Body    string
}

func (m *Model) CreateNewModel(ctx context.Context) (int64, error) {
  query := "INSERT INTO models (title, body) VALUES (?, ?);"
  stmt, err := db.DB.Prepare(query)
  if err != nil {
    return 0, err
  }

  defer stmt.Close()

  result, err := stmt.ExecContext(ctx, m.Title, m.Body)
  if err != nil {
    return 0, err
  }

  id, err := result.LastInsertId()
  if err != nil {
    return 0, err
  }
   
  return id, nil
}

func GetAllModels(ctx context.Context) ([]Model, error) {
  query := "SELECT id, title, body FROM models;"
  rows, err := db.DB.QueryContext(ctx, query)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  models := []Model{}
  for rows.Next() {
    model := Model{}
    err := rows.Scan(&model.Id, &model.Title, &model.Body)
    if err != nil {
      return nil, err
    }
    
    models = append(models, model)
  }
  
  return models, nil
}
`, moduleName)

  modelFile, err := os.OpenFile(modelFilePath, os.O_WRONLY|os.O_APPEND, 0666)
  if err != nil {
    return err
  }
  defer modelFile.Close()

  // write in model.go
  _, err = modelFile.Write([]byte(packageContent))
  if err != nil {
    return err
  }

  return nil
}
