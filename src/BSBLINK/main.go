package main

import (
	"bsblink/github.com/gorilla/mux"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello world")
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	http.ListenAndServe(":9090", router)
}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println("ada yang akses jam ", timeString)
}
