package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

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
	http.ListenAndServe(":9092", router)

}

func login(w http.ResponseWriter, r *http.Request) {
	//read response ===========================================================
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
	//read response ===========================================================

	//hit ke API KAK ANDY ===========================================================

	log.Println("Hit ke 172.17.34.58:9000")

	url := "http://172.17.34.58:9000/auth"

	// Metode HTTP
	method := "POST"

	// Data yang akan dikirim dalam body (jika metode adalah POST, PUT, PATCH, dll.)
	// payload := strings.NewReader("key1=value1&key2=value2")
	payload := strings.NewReader(contentString)
	// Membuat objek Request
	// req, err := http.NewRequest(method, url, payload)
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Menambahkan header ke permintaan (opsional)
	req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("Authorization", "Bearer YOUR_ACCESS_TOKEN")

	// Melakukan permintaan HTTP
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	//fmt.Println("Response Status:", resp.Body)

	// Menampilkan hasil respons
	fmt.Println("Response Status:", resp.Status)
	// ... (mengolah body respons atau melakukan operasi lain sesuai kebutuhan
	// Read the content of the response body into a byte slice
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Convert the byte slice to a string
	bodyString := string(body)

	// Print or use the string as needed
	fmt.Println("Respon Body : ", bodyString)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]byte(bodyString))
	log.Println("haaa")
}

func InqSaldo(w http.ResponseWriter, r *http.Request) {

}
