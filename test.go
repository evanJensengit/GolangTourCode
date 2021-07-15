package main

import (
	"fmt"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {
	a := Page{
		Title: "Test Page",
		Body:  []byte("This is a test page"),
	}

	fmt.Println(a.Title, a.Body)
	fmt.Printf("%s", a.Body)

	slice1 := make([]Page, 10)
	slice1[0] = Page {
		Title: "A story about Evan",
		Body: []byte("Evan's story began on a rainy day in Redmond..."),
	}
}


