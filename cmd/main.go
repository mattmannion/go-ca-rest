package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", GetPosts).Methods(http.MethodGet)
	r.HandleFunc("/", PostPost).Methods(http.MethodPost)

	fmt.Println("server live")
	log.Fatalln(http.ListenAndServe(":7890", r))
}
