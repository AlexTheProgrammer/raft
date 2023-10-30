package main

import (
	"log"
	"net/http"
)

func main() {
	port := ":3000"
	log.Printf("listening on port %v", port)
	log.Fatal(http.ListenAndServe(port, http.FileServer(http.Dir("../../assets"))))
}
