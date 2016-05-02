package main

import (
	"net/http"
)

func routes() {
	http.HandleFunc("/", index)
}

func main() {
	routes()
	http.ListenAndServe(":8080", nil)
}
