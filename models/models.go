package models

import "os"

type Page struct {
	Title string
	Body []byte
	FileList []string
}

var path string = "./files/"

func (p *Page) save() error {
	filename := path + p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}


type AuthData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


type RegisterData struct {
    Username       string `json:"username"`
    Password       string `json:"password"`
    Email          string `json:"email"`
    FavoritePokemon string `json:"favoritePokemon"`
}
