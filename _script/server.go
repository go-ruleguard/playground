package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("Listening on :8080...")
	log.Println(http.ListenAndServe(":8080", nil))
}
