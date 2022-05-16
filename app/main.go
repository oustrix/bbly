package main

import (
	"bbly/internal/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/", handlers.IndexHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
