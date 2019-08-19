// Basic Wiki Support Code

package main

import "io/ioutil"

// A Wiki Page
type Page struct {
	Title string
	Body  []byte
}

// When we want to save a page, the title becomes the name of
// the text file we save it to
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// Given the name of a page, we load the file, but if the
// page does not exist, then we return nil

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
