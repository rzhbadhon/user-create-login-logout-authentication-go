package handlers

import (
	"encoding/json"
	"go-auth-manual/models"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Handler struct {
	DB       *sqlx.DB
	Validate *validator.Validate
}

func NewHandler(db *sqlx.DB) *Handler {
	return &Handler{
		DB:       db,
		Validate: validator.New(),
	}
}

func (h *Handler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	// creating new user struct
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// validating the struct if all required is present 
	err = h.Validate.Struct(user)
	if err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	// adding other fields manually
	user.ID = uuid.New() // unique id
	user.Role = "user"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// database parsing
	query := `
	INSERT INTO users (id, first_name, last_name, email, password, role, created_at, updated_at)
	VALUES (:id, :first_name, :last_name, :email, :password, :role, :created_at, :updated_at)
	`

	_, err = h.DB.NamedExec(query, &user)
	if err != nil{
		pqErr, ok := err.(*pq.Error)
		if ok && pqErr.Code == "23505"{
			http.Error(w, "Email already exist", http.StatusConflict)
			return
		}
		log.Println("Database insert error: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
