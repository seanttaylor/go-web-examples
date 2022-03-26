package main

import (
	"html/template"
	"net/http"
)

func main() {
	type Todo struct {
		Title string
		Done  bool
	}

	type TodoPageData struct {
		PageTitle string
		Todos     []Todo
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("layout.html"))
		data := TodoPageData{
			PageTitle: "My Todo List",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}

		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":80", nil)
}
