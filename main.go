package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main(){

	Key := GenerateRandomKey()
	SetJWTKey(Key)
	r := gin.Default()

	SetupRoutes(r)

	r.Run(":"+port)
	log.Println("Server is running on port: ", port)
}