package main

import (
	"log"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseFiles("edit.html", "view.html", "list.html"))

func main() {
	log.Print("Starting to listen on port 8888")
    http.HandleFunc("/view/", makeHandler(viewHandler))
    http.HandleFunc("/edit/", makeHandler(editHandler))
    http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/delete/", makeHandler(deleteHandler))
	http.HandleFunc("/list/", listHandler)
    log.Fatal(http.ListenAndServe(":8888", nil))
}
