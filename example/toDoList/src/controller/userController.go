package main

import (
	"encoding/json"
	//"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Message struct {
	Content string `json:"content"`
	Id      int    `json:"id"`
}

func main() {

	//newMessage := Message{Content: "Hello World!", Id: 23}
	//
	//byteArray, err := json.Marshal(newMessage)
	//if err != nil {
	//	log.Fatal(err)
	//}

	mux := http.NewServeMux()
	PORT := ":8080"

	mux.Handle("/", http.HandlerFunc(getMessage))

	server := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	serverError := server.ListenAndServe()
	if serverError != nil {
		log.Fatal(server)
	}

}

func getMessage(response http.ResponseWriter, request *http.Request) {
	log.Println("the request is", request.URL.Path, "from", request.Host)
	response.WriteHeader(http.StatusOK)
	body := Message{Content: "Hello World", Id: 20}
	jsonString, err := json.Marshal(body)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Fprintf(response, "%s", jsonString)

}

//func defaultHandler(response http.ResponseWriter, request *http.Request) {
//	log.Println("default", request.URL.Path, "from", request.Host)
//	response.WriteHeader(http.StatusNotFound)
//	body := "thank you: try again later"
//	fmt.Fprintf(response, "%s", body)
//
//}
//
//func deleteHandler(response http.ResponseWriter, request *http.Request) {
//	parameter := strings.Split(request.URL.Path, "/")
//	fmt.Println(parameter)
//	response.WriteHeader(http.StatusOK)
//	if len(parameter) < 4 {
//		response.WriteHeader(http.StatusBadRequest)
//
//	}
//}
