package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/udborets/chat-app/server/internal/auth"
)

func StartServer() {
	// config for database
	config := "host=localhost port=5432 user=postgres password=1234 dbname=chat-app sslmode=disable"

	app := auth.NewHTTP(config)

	router := gin.Default()
	app.InitRoutes(router)

	router.Run(":1773")
}
