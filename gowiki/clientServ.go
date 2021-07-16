package main

//simple web server
//tutorial from https://golang.org/doc/articles/wiki/
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"html/template"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

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

//http.ResponseWriter handles assembles http servers response and by writing to it we send data to http client. http.Request is data struc that represtents client http request
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:]) //we print what is after "w," to the w and that becomes the server response
}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi this is handler1 saying fuck off %s!", r.URL.Path[1:]) //we print what is after "w," to the w and that becomes the server response
}

func viewHandler(w http.ResponseWriter, r *http.Request) { 
	title, err := getTitle(w,r)
	if err != nil {
		return
	}
	p, err := loadPage(title) 
	if err != nil { 
		http.Redirect(w,r, "/edit/" +title, http.StatusFound) 
		return
	}
 	renderTemplate(w, "view", p) }

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

func saveHandler(w http.ResponseWriter, r *http.Request) {
    title, err := getTitle(w, r)
	if err != nil {
		return
	}
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
  	err :=  p.save()
   	if err != nil {
   		http.Error(w, err.Error(), http.StatusInternalServerError)
   		return
   	}
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl. "html", p)
    //reads content of html file and returns a *template.Template
   
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    t.Execute(w, p) //this writes the generated html to the http.REsponseWriter
    if err != nill {
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

func main() { 
	//the "/" tells the http package to handle all requests to web root /" with the function "handler"

	http.HandleFunc("/", handler)       //HandleFunc is a method build into http imported thing
	http.HandleFunc("/balls", handler1) //HandleFunc is a method build into http imported thing
	http.HandleFunc("/view/", viewHandler)

	http.HandleFunc("/edit/", editHandler)
    http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil)) //log.Fatal used to record what is returned from listenandserve.

	//the 8080 says "listen on this port"
}
