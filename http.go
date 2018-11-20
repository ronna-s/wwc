package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Women Who Code Berin!"))
	})

	http.ListenAndServe(":8080", nil)
}
