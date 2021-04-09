package main

import (
	"fmt"
	"log"
	"net/http"

	handler "go-rest/pkg/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/", handler.SayHello("Youtube")).Methods("GET")
	r.HandleFunc("/user", handler.HandleUser).Methods("POST")

	// Server config
	fmt.Println("Server at PORT 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
