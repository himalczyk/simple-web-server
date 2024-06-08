package main

import (
	"log"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseFiles("templates/edit.html", "templates/view.html", "templates/list.html", "templates/auth.html"))

func main() {
	log.Print("Starting to listen on port 8888")
	http.HandleFunc("/auth/", authHandler)
	http.HandleFunc("/login/", loginHandler)
    http.HandleFunc("/view/", makeHandler(viewHandler))
    http.HandleFunc("/edit/", makeHandler(editHandler))
    http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/delete/", makeHandler(deleteHandler))
	http.HandleFunc("/list/", listHandler)
    log.Fatal(http.ListenAndServe(":8888", nil))
}
