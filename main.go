package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/himalczyk/simple-web-server/db"
)

var templates = template.Must(template.ParseFiles("templates/edit.html", "templates/view.html", "templates/list.html", "templates/auth.html", "templates/register.html", "templates/index.html"))

func main() {
	ctx := context.Background()
	dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")

    connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
    client, err := db.NewClient(ctx, connStr)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer client.Close()

	log.Print("Starting to listen on port 8888")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/auth/", authHandler)
	http.HandleFunc("/login/", loginHandler)
	http.HandleFunc("/registration/", registerHandler)
	http.HandleFunc("/register/", registerProcessHandler)
    http.HandleFunc("/view/", makeHandler(viewHandler))
    http.HandleFunc("/edit/", makeHandler(editHandler))
    http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/delete/", makeHandler(deleteHandler))
	http.HandleFunc("/list/", listHandler)
    log.Fatal(http.ListenAndServe(":8888", nil))
}
