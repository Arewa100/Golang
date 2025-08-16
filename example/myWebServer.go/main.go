package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("building my first web server in go")
	PORT := ":8080"

	mux := http.NewServeMux()
	fmt.Println("Listening on port", PORT)
	mux.Handle("/", http.HandlerFunc(handlerFunc))
	err := http.ListenAndServe(PORT, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func handlerFunc(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello Go! this is working")
	fmt.Println(request.Host)
}
