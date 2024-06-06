package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"regexp"
)

var validPath = regexp.MustCompile("^/(edit|save|view|delete)/([a-zA-Z0-9]+)$")

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	log.Printf("Loaded file %s", filename)
	return &Page{Title: title, Body: body}, nil
}


func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	defer log.Println("Rendered page", tmpl)
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
    m := validPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
        http.NotFound(w, r)
        return "", errors.New("invalid Page Title")
    }
    return m[2], nil // The title is the second subexpression.
}


func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
    if err != nil {
        return
    }
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusNotFound)
		return
	}
	renderTemplate(w, "view", p)
}


func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
    if err != nil {
        return
    }
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err = p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
    if err != nil {
        return
    }
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}


func deleteHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		http.Error(w, "Invalid file title", http.StatusBadRequest)
		return
	}
	err = delete(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusNotFound)
}
