package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
	Jam     string `json:"jam"`
}

func main() {
	log.Println("Halo pajri")

	router := mux.NewRouter()
	router.HandleFunc("/halo", Halo).Methods("GET")
	http.ListenAndServe(":9090", router)

}

func Halo(w http.ResponseWriter, r *http.Request) {

	now := time.Now()

	// Mengonversi waktu ke string dengan format default
	timeString := now.String()

	response := Response{Message: "Halo, Pajri!", Jam: timeString}

	log.Println("ada yang akses jam ", timeString)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
