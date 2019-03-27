package go_httpserver_sample

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func indexHandler(w http.ResponseWriter,
	r *http.Request) {

	fmt.Fprint(w, "hello world")
}

func personHandler(w http.ResponseWriter,
	r *http.Request) {
	defer r.Body.Close()

	if r.Method == "POST" {
		// Translate request body to Json
		var person Person
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}

		filename := fmt.Sprintf("%d.txt", person.ID)
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// Write name to the file
		_, err = file.WriteString(person.Name)
		if err != nil {
			log.Fatal(err)
		}

		// Send status code 201 as a response
		w.WriteHeader(http.StatusCreated)
	}
}

func Start() {
	fmt.Print("Open http://localhost:3000/")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/persons", personHandler)
	http.ListenAndServe(":3000", nil)
}
