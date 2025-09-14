package service

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
)

type Page struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

var validTitle = regexp.MustCompile("^[a-zA-Z0-9]+$") //合法title仅含字母数字

func (p *Page) Save() error {
	if !validTitle.MatchString(p.Title) {
		return errors.New("invalid title: only letters and numbers are allowed")
	}
	filename := p.Title + ".txt"
	fileStoragePath := filepath.Join("data", filename)
	return os.WriteFile(fileStoragePath, []byte(p.Body), 0600)
}
func LoadPage(title string) (*Page, error) {
	if !validTitle.MatchString(title) {
		return nil, errors.New("invalid title: only letters and numbers are allowed")
	}
	filename := title + ".txt"
	fileStoragePath := filepath.Join("data", filename)
	bodyBytes, err := os.ReadFile(fileStoragePath)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: string(bodyBytes)}, nil
}
