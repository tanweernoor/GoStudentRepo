// Basic Wiki Server Version 2

package main

import (
	"html/template"
	"net/http"
)

// The handler to manage a view request.
func viewHandler(wr http.ResponseWriter, req *http.Request) {
	// Extract the title and fetch the Wiki Page
	title := req.URL.Path[len("/view/"):]
	pp, err := loadPage(title)
	// if the page does not exist, then redirect to the edit so
	// that one can be added.
	if err != nil {
		http.Redirect(wr, req, "/edit/"+title, http.StatusFound)
		return
	}
	// Create a template and populate it with values
	t, _ := template.ParseFiles("view.html")
	t.Execute(wr, pp)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the page title from the URL
	title := r.URL.Path[len("/edit/"):]
	// Load the page, and if it cannot be loaded, then
	// provide a blank text area with the supplied title.
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	// Create a template and populate it
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	// We extract the value of the textarea and put it into
	// a wiki page and save it.
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	// Once saved, the file is then opened for display
	// by the /view/handler
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
func main() {
	// The two handlers implemented so far
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":3001", nil)
}
