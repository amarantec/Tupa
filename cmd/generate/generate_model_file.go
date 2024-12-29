package generate

import (
	"fmt"
	"os"
)

func WriteModelFile(modelFilePath, moduleName string) error {
	/*
	   Code to write in model.go
	*/

	packageContent := fmt.Sprintf(`package internal

import (
  "context"
  "%s/internal/db"
)

type %s struct {
  Id      int64
  Title   string   
  Body    string
}

func (m *%s) Create(ctx context.Context) (int64, error) {
  query := "INSERT INTO %ss (title, body) VALUES (?, ?);"
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

func GetAll(ctx context.Context) ([]%s, error) {
  query := "SELECT id, title, body FROM %ss;"
  rows, err := db.DB.QueryContext(ctx, query)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  %ss := []%s{}
  for rows.Next() {
    %s := %s{}
    err := rows.Scan(&%s.Id, &%s.Title, &%s.Body)
    if err != nil {
      return nil, err
    }
    
    %ss = append(%ss, %s)
  }
  
  return %ss, nil
}
`, moduleName, moduleName, moduleName, moduleName, moduleName,
		moduleName, moduleName, moduleName, moduleName, moduleName, moduleName,
		moduleName, moduleName, moduleName, moduleName, moduleName, moduleName)

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
