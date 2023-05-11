package repository

import (
	"database/sql"
	"log"
)

var database *sql.DB

func InitDB(config string) {
	db, err := sql.Open("postgres", config)
	if err != nil {
		log.Fatalf(err.Error())
	}
	database = db
}
