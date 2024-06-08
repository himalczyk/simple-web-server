package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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
		if strings.Contains(err.Error(), "no template") {
			errorString := fmt.Sprintf("Template: %v not registered in templates %v", tmpl+".html", templates.DefinedTemplates())
			http.Error(w, errorString, http.StatusInternalServerError)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
    }
}


func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        m := validPath.FindStringSubmatch(r.URL.Path)
        if m == nil {
            http.NotFound(w, r)
            return
        }
        fn(w, r, m[2])
    }
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusNotFound)
		return
	}
	renderTemplate(w, "view", p)
}


func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}


func deleteHandler(w http.ResponseWriter, r *http.Request, title string) {
	err := delete(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusNotFound)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	txtFiles, err := findTxtFiles(".")
	if err != nil {
		return
	}
	p := &Page{FileList: txtFiles}
	log.Println(txtFiles)
	renderTemplate(w, "list", p)
}

func findTxtFiles(dir string) ([]string, error) {
	var txtFiles []string

	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(d.Name()) == ".txt" {
			fileName := strings.Split(path, ".")[0]
			txtFiles = append(txtFiles, fileName)
		}
		return nil
	})

	return txtFiles, err
}


func authHandler(w http.ResponseWriter, r *http.Request) {
	p := &Page{Title: "Login"}
	renderTemplate(w, "auth", p)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// handle auth here as the coming in username in password is here
    // Parse form data
    username := r.FormValue("username")
    password := r.FormValue("password")

	// add here checking in db if account exists and his password is correct

    // For demonstration, we'll just print the credentials to the console
    // In a real application, you should verify the credentials
    fmt.Printf("Username: %s, Password: %s\n", username, password)

    // Redirect to a success page or display a message
    fmt.Fprintf(w, "Login successful!")
}
