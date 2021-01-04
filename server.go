package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

type Mahasiswa struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

// merubah ke JSON
func encodeJson(data []Mahasiswa) []byte {
	dataJson, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	return dataJson
}

// routing
func Home(w http.ResponseWriter, r *http.Request) {
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
}

func getMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		mhs := []Mahasiswa{
			{
				1,
				"Walid nurudin",
				"Indramayu",
			},
			{
				2,
				"Khairun arkham",
				"Karanganyar",
			},
			{
				3,
				"Jembar ashofa",
				"Cirebon",
			},
			{
				4,
				"Muhammad zulfadli",
				"Medan",
			},
		}
		dataMahasiswa := encodeJson(mhs)

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataMahasiswa)
		return
	}

	http.Error(w, "Hayo mau ngapain", http.StatusNotFound)

}

func main() {
	// Static file
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	// endpoint
	http.HandleFunc("/", Home)
	http.HandleFunc("/mahasiswa", getMahasiswa)

	fmt.Println("starting web server at http://localhost:3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
