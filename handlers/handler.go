package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
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
