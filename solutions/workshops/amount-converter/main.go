package main

import (
	"log"
	"net/http"

	"github.com/aubm/golang-codelab/solutions/workshops/amount-converter/amounts"
	"github.com/aubm/golang-codelab/solutions/workshops/amount-converter/api"
)

func main() {
	converter := &amounts.Converter{FixerApiUrl: "http://api.fixer.io"}
	amountHandlers := &api.AmountHandlers{Converter: converter}

	http.HandleFunc("/convert", amountHandlers.Convert)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
