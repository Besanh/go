package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os" // lam viec voi file
	"regexp"
	"text/template"
)

// Store in memory
type Page struct {
	Title string
	Body  []byte
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))

	/**
	* HandleFunc noi http biet de handle tat ca request toi web root("/") thong qua hanlder()
	* ListenAndServe luon return 1 error,
	 */
	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHanlder)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

/**
* p la con tro
* save() nhan p la con tro toi Page,
* khong co tham so, return gia tri loi
* Neu khong loi, Page.save() return @nil
* 0600 la permission(read,write)
 */
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

/**
* ReadFile return []byte va error
* Dau "_" de loai bo gia tri tra ve loi(chi la phep gan bt)
 */
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// Func nay la 1 dang cua http.HandleFunc
/**
* http.ResponseWriter tra ve response cua HTTP server, sau do gui data toi HTTP client
* http.Request la 1 cau truc data dai dien cho HTTP client
* r.URL.Path la url yeu cau, [1:] loai bo dau "/" dau tien ra khoi path
 */
func handler(w http.ResponseWriter, r *http.Request) {
	// Open va write vao w
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func viewHanlder(w http.ResponseWriter, r *http.Request) {
	// title := r.URL.Path[len("/view/"):]
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusNotFound)
		return
	}
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	// t, _ := template.ParseFiles("view.html")
	// t.Execute(w, p)
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	// title := r.URL.Path[len("/edit/"):]
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	// fmt.Fprintf(w, "<h1>Editing %s</h1>"+
	// 	"<form action=\"/save/%s\" method=\"POST\">"+
	// 	"<textarea name=\"body\">%s</textarea><br>"+
	// 	"<input type=\"submit\" value=\"Save\">"+
	// 	"</form>", p.Title, p.Title, p.Body)

	// @return a*template.Template
	// t, _ := template.ParseFiles("edit.html")
	// Xuat html vao ResponseWriter
	// t.Execute(w, p)

	renderTemplate(w, "edit", p)

}

// FormValue la type string
func saveHandler(w http.ResponseWriter, r *http.Request) {
	// title := r.URL.Path[len("/save/"):]
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// Render file html hoac return error
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalidate page title")
	}
	return m[2], nil // The title is the second subexpression.
}
