package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
