package main

import (
	"log"
	"os"
)

type Page struct {
	Title string
	Body []byte
	FileList []string
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func delete(title string) error {
	filename := path + title + ".txt"
	defer log.Printf("Deleted file %s", filename)
	return os.Remove(filename)
}
