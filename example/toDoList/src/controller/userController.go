package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	//"toDoList/src/services/"
)

func main() {
	PORT := ":8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/register", http.HandlerFunc(RegisterUser))

	server := http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}

	severError := server.ListenAndServe()
	if severError != nil {
		log.Fatal("this is the server error: ", severError)
	}
}

func RegisterUser(response http.ResponseWriter, request *http.Request) {
	fmt.Println("SERVING", request.URL.Path, "from", request.Host)
	if request.Method != "POST" { //check request method
		http.Error(response, "method not allowed", http.StatusMethodNotAllowed)
		response.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	userIncomingData, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(response, "bad request", http.StatusBadRequest)
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	var userData []struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err = json.Unmarshal(userIncomingData, &userData)
	fmt.Println(userIncomingData)

}
