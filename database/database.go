package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func ConnectDB() *sqlx.DB{
	connStr := "user=postgres password:1212 dbname=yourdb sslmode=disable"

	db ,err := sqlx.Connect("psotgres", connStr)
	if err != nil{
		log.Fatal("failed to connect to database: ", err)
	}

	log.Println("Database connected succssfully")
	return db
}