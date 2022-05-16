package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
