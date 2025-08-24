package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name    string
	Age     int
	Country string
}

var newPerson Person = Person{"Miracle", 25, "Nigeria"}
var newJsonData, err = json.Marshal(newPerson)

func main() {
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(newJsonData))
	fmt.Println("quick test of my application")
	PORT := ":8080"

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(PrintName))

	serverError := http.ListenAndServe(PORT, mux)
	if serverError != nil {
		log.Fatal(serverError)
	}

}

func PrintName(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "<h1>Hello World<h1/>", newJsonData)
}
