package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to homepage")
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only post method is required", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Println("Error request body: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(user)
	if err != nil{
		log.Println("Error encoding JSON: ", err)
		return
	}

	fmt.Fprintln(w, "This is signup page (POST)")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only post method is required", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "This is login page (POST)")
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/signup", SignupHandler)
	mux.HandleFunc("/login", LoginHandler)

	fmt.Println("Server is starting on port :9000...")

	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		log.Println("Error starting server ", err)
	}
}
