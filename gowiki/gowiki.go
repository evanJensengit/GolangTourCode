package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600) //this writes the body to a file and creates the file if it is not there, the "0600" is code for read write
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename) //gets the body from the file with the name of "filename"
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil //returns page reference with title and body
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")} //create page object
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
