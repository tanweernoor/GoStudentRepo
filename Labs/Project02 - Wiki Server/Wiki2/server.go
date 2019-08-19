// Basic Wiki Server

package main

import (
	"html/template"
	"net/http"
)

// The handler to manage a view request.
func viewHandler(wr http.ResponseWriter, req *http.Request) {
	// Extract the title and fetch the Wiki Page
	title := req.URL.Path[len("/view/"):]
	pp, _ := loadPage(title)
	// Create a template and populate it with values
	t, _ := template.ParseFiles("view.html")
	t.Execute(wr, pp)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the page title from the URL
	title := r.URL.Path[len("/edit/"):]
	// Load the page, and if it cannot be loaded, then
	// provide a blank text area withe supplied title.
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	// Create a template and populate it
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
}

func main() {
	// The two handlers implemented so far
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":3001", nil)
}
