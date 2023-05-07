package internal

import "github.com/IskanderSh/chat-app/internal/delivery"

func StartServer() {
	// config for database

	app := delivery.NewHTTP()
	app.Start()
}
