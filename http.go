package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hi/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello, Women Who Code Berlin!</h1>"))
	})

	http.ListenAndServe(":8080", nil)
}
