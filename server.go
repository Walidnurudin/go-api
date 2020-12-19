package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func main() {
	// Static file
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	// Routing
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name": "Walid nurudin",
		}

		var filepath = path.Join("views", "index.html")
		var te, err = template.ParseFiles(filepath)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// Execute() akan membuat hasil parsing template ditampilkan ke layar web browser.
		te.Execute(w, data)
	})

	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
