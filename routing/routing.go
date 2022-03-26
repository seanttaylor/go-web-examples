// https://gowebexamples.com/routes-using-gorilla-mux/
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books/{title}/pages/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"] // the book title
		page := vars["page"]   // the page number

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)

	})
	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to index route")
	})*/

	http.ListenAndServe(":80", router)
}
