package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/himalczyk/simple-web-server/models"
)

var validPath = regexp.MustCompile("^/(edit|save|view|delete)/([a-zA-Z0-9]+)$")

var path string = "./files/"

func delete(title string) error {
	filename := path + title + ".txt"
	defer log.Printf("Deleted file %s", filename)
	return os.Remove(filename)
}

func loadPage(title string) (*models.Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(path + filename)
	if err != nil {
		return nil, err
	}
	log.Printf("Loaded file %s", filename)
	return &models.Page{Title: title, Body: body}, nil
}


func renderTemplate(w http.ResponseWriter, tmpl string, p *models.Page) {
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

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := &models.Page{Title: "Index"}
	renderTemplate(w, "index", p)
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
	p := &models.Page{Title: title, Body: []byte(body)}
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
		p = &models.Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}


func deleteHandler(w http.ResponseWriter, r *http.Request, title string) {
	err := delete(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/list/", http.StatusOK)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	txtFiles, err := findTxtFiles(".")
	if err != nil {
		return
	}
	p := &models.Page{FileList: txtFiles}
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
			fileName = strings.Split(fileName, "/")[1]
			txtFiles = append(txtFiles, fileName)
		}
		return nil
	})

	return txtFiles, err
}


func authHandler(w http.ResponseWriter, r *http.Request) {
	p := &models.Page{Title: "Login"}
	renderTemplate(w, "auth", p)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	AuthData := &models.AuthData{
        Username:       r.FormValue("username"),
        Password:       r.FormValue("password"),
    }
	log.Println(AuthData)

	// add here checking in db if account exists and his password is correct
    // Redirect to a success page or display a message
	// errors etc.
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	p := &models.Page{Title: "Register"}
	renderTemplate(w, "register", p)
}

func registerProcessHandler(w http.ResponseWriter, r *http.Request) {
	registerData := &models.RegisterData{
        Username:       r.FormValue("username"),
        Password:       r.FormValue("password"),
        Email:          r.FormValue("email"),
        FavoritePokemon: r.FormValue("favoritePokemon"),
    }
	log.Println(registerData)
	// add here saving to db for account creation
	// Redirect to a success page, errors etc.
	fmt.Fprintf(w, "Form Data: %+v\n", registerData)
	fmt.Fprintf(w, "Register successful!")
}
