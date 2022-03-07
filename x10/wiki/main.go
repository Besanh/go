package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"text/template"
)

type Page struct {
	Title string
	Body  []byte
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func main() {
	// Gan p1 vao dia chi Page
	p1 := &Page{Title: "AnhLe", Body: []byte("<?php echo 'Handsome' ?>")}
	// Vi trong ham seen() co p *Page nen p1 se return gia tri cuar seen()
	p1.seen()
	p2, _ := loadPage(p1.Title)
	fmt.Println(string(p2.Body))

	http.HandleFunc("/", homePage)
	http.HandleFunc("/view/", makeHandler(viewPage))
	http.HandleFunc("/edit/", makeHandler(editPage))
	http.HandleFunc("/save/", makeHandler(savePage))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Tao file
// Func nay ten la seen
// Cach dung con tro nhung khong muon con tro la argument
// Tra ve gia tri cua con tro
func (p *Page) seen() error {
	filename := p.Title + ".php"
	return os.WriteFile(filename, p.Body, 0600)
}

// Load page
func loadPage(name string) (*Page, error) {
	filename := name + ".php"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err // vi func return 2 gia tri
	}
	return &Page{Title: name, Body: body}, nil
}

func homePage(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hi you, %v\n", request.URL.Path[1:])
}

func viewPage(response http.ResponseWriter, request *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(response, request, "/edit/"+title, http.StatusNotFound)
		return
	}
	renderTemplate(response, "view", p)
}

func editPage(response http.ResponseWriter, request *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(response, "edit", p)
}

func savePage(response http.ResponseWriter, request *http.Request, title string) {
	body := request.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.seen()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(response, request, "/view/"+title, http.StatusFound)
}

// Doc du lieu tu path
func getTitle(response http.ResponseWriter, request *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(request.URL.Path)
	if m == nil {
		http.NotFound(response, request)
		return "", errors.New("Invalidate page title")
	}
	return m[2], nil // The title is the second subexpression.
}

// Render template
func renderTemplate(response http.ResponseWriter, name string, p *Page) {
	err := templates.ExecuteTemplate(response, name+".html", p)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

/**
* Function closure
* @return cua func makeHandle la 1 func nhan ResponseWriter va Request(2 thang nay thuc chat la HandleFunc)
 */
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		// Here we will extract the page title from the Request,
		// and call the provided handler 'fn'
		m := validPath.FindStringSubmatch(request.URL.Path)
		if m == nil {
			http.NotFound(response, request)
			return
		}
		fn(response, request, m[2])
	}
}
