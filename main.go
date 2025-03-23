package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello world")

	handler_one := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))

		films := map[string][]Film{
			"Films": {
				{Title: "The GodFather", Director: "Francis Ford Coppola"},
				{Title: "Pulp Fiction", Director: "Quentin Tarantino"},
				{Title: "Inception", Director: "Christopher Nolan"},
				{Title: "The Dark Knight", Director: "Christopher Nolan"},
				{Title: "Fight Club", Director: "David Fincher"},
				{Title: "Forrest Gump", Director: "Robert Zemeckis"},
				{Title: "The Matrix", Director: "Lana Wachowski, Lilly Wachowski"},
				{Title: "The Shawshank Redemption", Director: "Frank Darabont"},
				{Title: "The Lord of the Rings: The Return of the King", Director: "Peter Jackson"},
				{Title: "Star Wars: Episode IV - A New Hope", Director: "George Lucas"},
				{Title: "The Silence of the Lambs", Director: "Jonathan Demme"},
			},
		}

		templ.Execute(w, films)
	}

	http.HandleFunc("/", handler_one)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
