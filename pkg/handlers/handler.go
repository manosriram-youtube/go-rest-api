package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GoUser struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func HandleUser(w http.ResponseWriter, r *http.Request) {
	var user GoUser
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	json.NewEncoder(w).Encode(user)
}

func SayHello(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Hello %s", name)
	}
}
