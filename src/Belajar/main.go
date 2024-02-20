package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type LoginResponse struct {
	Name string `json:"name"`
}

type Body struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	http.ListenAndServe(":9090", router)

}

func login(w http.ResponseWriter, r *http.Request) {
	var L LoginResponse
	L.Name = "TACA"
	log.Println("ada yang akses")

	log.Println("ISI :", r.Body)

	// Read the content of the io.ReadCloser into a byte slice
	content, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the byte slice to a string
	contentString := string(content)

	// Print or use the string as needed
	log.Println(contentString)

	//-----------------------------------------------------------
	var b Body
	err1 := json.Unmarshal([]byte(content), &b)
	if err1 != nil {
		fmt.Println("Error unmarshaling JSON:", err1)
		return
	}

	log.Println("\nIsi Username : ", b.Username)
	log.Println("Isi Password : ", b.Password)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(L)

}
