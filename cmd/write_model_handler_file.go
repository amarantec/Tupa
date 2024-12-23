package cmd

import (
	"fmt"
	"os"
)

func writeModelHandlerFile(modelHandlerFilePath, moduleName string) error {
  /* 
    Code to write in model_handler.go
  */

  packageContent := fmt.Sprintf(`package handler
  
import(
  "context"
  "encoding/json"
  "net/http"
  "time"
  "%s/internal"
)
  

func LoadCreateModelForm(w http.ResponseWriter, r *http.Request) {
  err := Templates.ExecuteTemplate(w, "create-model.html", nil)
  if err != nil {
    http.Error(w, "Error rendering template", http.StatusInternalServerError)
    return
  }
}
    
func CreateNewModel(w http.ResponseWriter, r *http.Request) {
  ctxTimeout, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  defer cancel()
  if r.Method == http.MethodPost {
    if r.Header.Get("Content-Type") == "application/json" {
      model := internal.Model{}
      if err :=
        json.NewDecoder(r.Body).Decode(&model); err != nil {
          http.Error(w, "could not decode this request. error: " + err.Error(), http.StatusBadRequest)
        return
      }


      response, err := model.CreateNewModel(ctxTimeout)
      if err != nil {
        http.Error(w, "could not create this model. error: " + err.Error(), http.StatusInternalServerError)
      }

      jsonResponse, _ := json.Marshal(response)
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(http.StatusCreated)
      w.Write(jsonResponse)

      return
    }
  r.ParseForm()

  title := r.FormValue("title")
  body := r.FormValue("body")

  newModel := internal.Model{Title: title, Body: body}
  newModel.CreateNewModel(ctxTimeout)

  http.Redirect(w, r, "/list", http.StatusSeeOther)
  } else {
    http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
  }
}

func GetAllModels(w http.ResponseWriter, r *http.Request) {
  ctxTimeout, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  defer cancel()
  
  response, err := internal.GetAllModels(ctxTimeout)
  if err != nil {
    http.Error(w, "could not get models. error: " + err.Error(), http.StatusInternalServerError)
    return
  }
  
  jsonResponse, _ := json.MarshalIndent(response, "", " ")

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(jsonResponse)
}

func ListModels(w http.ResponseWriter, r *http.Request) {
  ctxTimeout, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  defer cancel()

  tmpl := Templates.Lookup("list-model.html")
  if tmpl == nil {
    http.Error(w, "template not found", http.StatusInternalServerError)
    return
  }

  models, err := internal.GetAllModels(ctxTimeout)
  if err != nil {
    http.Error(w, "could not load models. error: " + err.Error(), http.StatusInternalServerError)
    return
  }

  err = tmpl.Execute(w, map[string]interface{}{
    "Model": models,
  })

  if err != nil {
    http.Error(w, "error rendering template", http.StatusInternalServerError)
    return
  }
}
`, moduleName)

  modelHandlerFile, err := os.OpenFile(modelHandlerFilePath, os.O_WRONLY|os.O_APPEND, 0666)
  if err != nil {
    return err
  }
  defer modelHandlerFile.Close()

  // write in model_handler.go
  _, err = modelHandlerFile.Write([]byte(packageContent))
  if err != nil {
    return err
  }
  return nil
}
