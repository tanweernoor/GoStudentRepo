// Basic Wiki Server Version 3

package main

import (
	"errors"
	"html/template"
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil
}

// The handler to manage a view request.
func viewHandler(wr http.ResponseWriter, req *http.Request) {
	// Extract the title and fetch the Wiki Page
	title, err := getTitle(wr, req)
	if err != nil {
		return
	}
	// if the page does not exist, then redirect to the edit so
	// that one can be added.
	pp, err := loadPage(title)
	if err != nil {
		http.Redirect(wr, req, "/edit/"+title, http.StatusFound)
		return
	}
	// Create a template and populate it with values
	t, _ := template.ParseFiles("view.html")
	t.Execute(wr, pp)
}

func editHandler(wr http.ResponseWriter, req *http.Request) {
	// Extract the page title from the URL
	title, err := getTitle(wr, req)
	if err != nil {
		return
	}
	// Load the page, and if it cannot be loaded, then
	// provide a blank text area with the supplied title.
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	// Create a template and populate it
	t, _ := template.ParseFiles("edit.html")
	t.Execute(wr, p)
}

func saveHandler(wr http.ResponseWriter, req *http.Request) {
	title, err := getTitle(wr, req)
	if err != nil {
		return
	}
	// We extract the value of the textarea and put it into
	// a wiki page and save it.
	body := req.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	// Once saved, the file is then opened for display
	// by the /view/handler
	http.Redirect(wr, req, "/view/"+title, http.StatusFound)
}
func main() {
	// The two handlers implemented so far
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":3001", nil)
}
