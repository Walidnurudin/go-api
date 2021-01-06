package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type student struct {
	Name  string `json:"name"`
	Grade int    `json:"Grade"`
}

// merubah ke JSON
func encodeJson(data []student) []byte {
	dataJson, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	return dataJson
}

func main() {
	var tmpl, err = template.ParseGlob("views/*")
	if err != nil {
		panic(err.Error())
		return
	}

	// Static file
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	// endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name": "Walid nurudin",
		}

		err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	fmt.Println("starting web server at http://localhost:3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
