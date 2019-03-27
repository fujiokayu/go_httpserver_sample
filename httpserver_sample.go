package go_httpserver_sample

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter,
	r *http.Request) {

	fmt.Fprint(w, "hello world")
}

func Start() {
	fmt.Print("Open http://localhost:3000/")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3000", nil)
}
