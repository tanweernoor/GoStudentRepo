// lab04-08.go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	response, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.Status)
	}
}
