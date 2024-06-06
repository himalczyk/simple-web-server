package main

import (
	"log"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func main() {
	log.Print("Starting to listen on port 8888")
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/delete/", deleteHandler)
    log.Fatal(http.ListenAndServe(":8888", nil))
}
