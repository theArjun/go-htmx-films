package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {

	home_page_handler := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("templates/index.html"))

		films := map[string][]Film{
			"Films": {
				{Title: "The GodFather", Director: "Francis Ford Coppola"},
				{Title: "Pulp Fiction", Director: "Quentin Tarantino"},
				{Title: "Inception", Director: "Christopher Nolan"},
				{Title: "The Dark Knight", Director: "Christopher Nolan"},
				{Title: "Forrest Gump", Director: "Robert Zemeckis"},
				{Title: "The Matrix", Director: "Lana Wachowski, Lilly Wachowski"},
				{Title: "The Shawshank Redemption", Director: "Frank Darabont"},
			},
		}

		templ.Execute(w, films)
	}

	add_film_handler := func(w http.ResponseWriter, r *http.Request) {
		// Simulate latency
		time.Sleep(1 * time.Second)

		// Extract POST Data
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		templ := template.Must(template.ParseFiles("templates/index.html"))
		templ.ExecuteTemplate(
			w,
			"film-list-element",
			Film{
				Title:    title,
				Director: director,
			},
		)
	}

	http.HandleFunc("/", home_page_handler)
	http.HandleFunc("/add-film/", add_film_handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
