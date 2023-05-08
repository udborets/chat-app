package internal

import (
	"github.com/udborets/chat-app/server/internal/delivery"
)

func StartServer() {
	// config for database
	config := "host=localhost port=5432 user=postgres password=1234 dbname=chat-app sslmode=disable"

	app := delivery.NewHTTP(config)
	app.Start()
}
