package generate

import (
	"fmt"
	"os"
)

func WriteModelHandlerFile(modelHandlerFilePath, moduleName string) error {
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
  
/*
  This function loads the web template to create a new model
*/

func LoadCreateModelTemplate(w http.ResponseWriter, r *http.Request) {
  err := Templates.ExecuteTemplate(w, "create-model-template.html", nil)
  if err != nil {
    http.Error(w, "Error rendering template", http.StatusInternalServerError)
    return
  }
}

/*
  This function makes the request in the API and then writes the database to create a new model
*/
    
func Create(w http.ResponseWriter, r *http.Request) {
  ctxTimeout, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  defer cancel()

  if r.Method == http.MethodPost {
    if r.Header.Get("Content-Type") == "application/json" {
      %s := internal.%s{}
      if err :=
        json.NewDecoder(r.Body).Decode(&%s); err != nil {
          http.Error(w, "could not decode this request. error: " + err.Error(), http.StatusBadRequest)
        return
      }

      response, err := %s.Create(ctxTimeout)
      if err != nil {
        http.Error(w, "could not create %s. error: " + err.Error(), http.StatusInternalServerError)
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

  new%s := internal.%s{Title: title, Body: body}
  new%s.Create(ctxTimeout)

  http.Redirect(w, r, "/list", http.StatusSeeOther)
  } else {
    http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
  }
}

/*
  This function makes a request to the API that lists all models available in the database
*/
func GetAll(w http.ResponseWriter, r *http.Request) {
  ctxTimeout, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  defer cancel()
  
  response, err := internal.GetAll(ctxTimeout)
  if err != nil {
    http.Error(w, "could not get %s. error: " + err.Error(), http.StatusInternalServerError)
    return
  }
  
  jsonResponse, _ := json.MarshalIndent(response, "", " ")

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(jsonResponse)
}

/*
  This function loads the template that lists all models available in the database  
*/
func ListModelsTemplate(w http.ResponseWriter, r *http.Request) {
  ctxTimeout, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  defer cancel()

  tmpl := Templates.Lookup("list-model-template.html")
  if tmpl == nil {
    http.Error(w, "template not found", http.StatusInternalServerError)
    return
  }

  %ss, err := internal.GetAll(ctxTimeout)
  if err != nil {
    http.Error(w, "could not load %ss. error: " + err.Error(), http.StatusInternalServerError)
    return
  }

  err = tmpl.Execute(w, map[string]interface{}{
    "%s": %ss,
  })

  if err != nil {
    http.Error(w, "error rendering template", http.StatusInternalServerError)
    return
  }
}
`, moduleName, moduleName, moduleName, moduleName,
		moduleName, moduleName, moduleName, moduleName, moduleName,
		moduleName, moduleName, moduleName, moduleName, moduleName)

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
