package models

import (
	"os"

	"gorm.io/gorm"
)

type Page struct {
	Title string
	Body []byte
	FileList []string
}

var path string = "./files/"

func (p *Page) Save() error {
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

type User struct {
	gorm.Model
	Username       string `gorm:"uniqueIndex;size:255"`
	Password       string
	Email          string `gorm:"uniqueIndex;size:255"`
	FavoritePokemon string
}
