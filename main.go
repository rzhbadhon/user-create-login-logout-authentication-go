package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name" validate: "required"`
	LastName  string    `json:"last_name"  validate:"required"`
	Email     string    `json:"email"    validate:"required,email"`
	Password  string    `json:"password"   validate:"required,min=6"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var validate = validator.New()

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

	err = validate.Struct(user)
	if err != nil{
		log.Println("Error hashing password", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil{
		log.Println("Error hashing password:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPassword)
	user.ID = ""

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(user)
	if err != nil{
		log.Println("Error encoding JSON: ", err)
		return
	}
	user.ID = "user_" + fmt.Sprintf("%d", time.Now().UnixNano())
	user.Role = "user"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	log.Printf("New user created successfully: %s (ID: %s)", user.Email, user.ID)

	user.Password = "" 

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println("Error encoding JSON response:", err)
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
