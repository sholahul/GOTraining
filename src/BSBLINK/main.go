package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	http.ListenAndServe(":9090", router)
}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println("ada yang akses")
}
