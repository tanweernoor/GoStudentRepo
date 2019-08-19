// Basic Wiki Server

package main

import (
	"fmt"
	"net/http"
)

// The handler to manage a view request.
func viewHandler(wr http.ResponseWriter, req *http.Request) {
	// The title of the page is the URL segment that appears after
	// the view request.
	title := req.URL.Path[len("/view/"):]
	// Load the page and write it to the response writer.
	pg, _ := loadPage(title)
	fmt.Fprintf(wr, "<h1>%s</h1><div>%s</div>", pg.Title, pg.Body)
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":3001", nil)
}
