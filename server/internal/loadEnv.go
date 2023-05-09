package internal

import (
	"github.com/joho/godotenv"
	"log"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
