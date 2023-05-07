package internal

import "github.com/udborets/chat-app/server/internal/delivery"

func StartServer() {
	// config for database
	config := "host=localhost user=postgres password=1234 dbname=test sslmode=disable"

	app := delivery.NewHTTP(config)
	app.Start()
}
