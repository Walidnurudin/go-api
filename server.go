package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name": "Walid nurudin",
		}

		var te, err = template.ParseFiles("template.html")
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
