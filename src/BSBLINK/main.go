package main

import (
<<<<<<< HEAD
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
=======
	"bsblink/controller"
	"bsblink/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type LoginResponse struct {
	Name string `json:"name"`
}

type Body struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	err := godotenv.Load(".env") //
	if err != nil {
		log.Println(err, ".env not load")
	}

	db := model.ConnectDB()
	defer db.Close()

	//create router
	router := mux.NewRouter()
	router.HandleFunc("/Auth", Auth).Methods("POST")
	router.HandleFunc("/GetUser", controller.GetUsers(db)).Methods("GET")
	router.HandleFunc("/GetUserByID/{id}", controller.GetUserByID(db)).Methods("GET")
	// router.HandleFunc("/Add", controller.CreateUser(db)).Methods("POST")

	http.ListenAndServe(":9090", router)

}

func Auth(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	// Mengonversi waktu ke string dengan format default
	timeString := now.String()
	log.Println("Ada yang hit Jam : ", timeString)

	// Read the content of the io.ReadCloser into a byte slice
	content, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the byte slice to a string
	contentString := string(content)

	// Print or use the string as needed
	log.Println(contentString)

	//CHECK USERNAME in body api
	var b Body
	err1 := json.Unmarshal([]byte(content), &b)
	if err1 != nil {
		fmt.Println("Error unmarshaling JSON:", err1)
		return
	}
>>>>>>> 99ba4ac2566d9071ba28e9df0815e8c23f6bcdaa
}
