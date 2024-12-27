package generate

const (
	mainFilePackageContentWithPostgres = `package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/joho/godotenv"
	"%s/internal/db"
	"%s/internal/handler/routes"
	"%s/internal/handler"
)

func main() {
	loadEnv()
	ctx := context.Background()

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatal("one or more environment variables are not set")
	}

	connectionString := fmt.Sprintf("host=%%s port=%%s user=%%s password=%%s dbname=%%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	DB, err := db.OpenConnection(ctx, connectionString)
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	handler.LoadTemplates()
  	mux := routes.SetRoutes()
  	mux.Handle("/css/",
    	http.StripPrefix("/css/",
      		http.FileServer(http.Dir("../../web/css"))))

	server := &http.Server{
		Addr:    		"127.0.0.1:8080",
		Handler: 		mux,
		ReadTimeout:  	10 * time.Second,
    	WriteTimeout: 	10 * time.Second,
	}

	fmt.Println("Server listen on: http://" + server.Addr)
	log.Fatal(server.ListenAndServe())

}

func loadEnv() {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal("error loading .env file")
	}
}`
	mainFilePackageContentWithSqlite = `package main

import ( 
  "fmt"
  "log"
  "net/http"
  "time"
  "%s/internal/handler"
  "%s/internal/handler/routes"
  "%s/internal/db"
)

func main() {
	db.OpenConnection()
	handler.LoadTemplates()
  	mux := routes.SetRoutes()
  	mux.Handle("/css/",
    	http.StripPrefix("/css/",
      		http.FileServer(http.Dir("../../web/css"))))

  	server := &http.Server{
    	Addr:         "127.0.0.1:8080",
    	Handler:      mux,
    	ReadTimeout:  10 * time.Second,
    	WriteTimeout: 10 * time.Second,
  	}

  	fmt.Println("Server listen on: http://" + server.Addr)
  	log.Fatal(server.ListenAndServe())
}`
	mainFilePackageContentWithMySql = `package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"%s/internal/db"
	"%s/internal/handler"
	"%s/internal/handler/routes"
	_ "github.com/go-sql-driver/mysql" // Driver MySQL
)

func main() {
	loadEnv()
	ctx := context.Background()

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatal("one or more environment variables are not set")
	}

	connectionString := fmt.Sprintf("%%s:%%s@tcp(%%s:%%s)/%%s",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	DB, err := db.OpenConnection(ctx, connectionString)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer DB.Close()

	handler.LoadTemplates()
	mux := routes.SetRoutes()
	mux.Handle("/css/",
		http.StripPrefix("/css/",
			http.FileServer(http.Dir("../../web/css"))))

	server := &http.Server{
		Addr:         "127.0.0.1:8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server listen on: http://" + server.Addr)
	log.Fatal(server.ListenAndServe())
}

func loadEnv() {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal("error loading .env file")
	}
}`
)
