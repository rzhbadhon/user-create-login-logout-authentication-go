package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fmt.Println("Server is starting on port :9000...")

	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		log.Println("Error starting server ", err)
	}
}
