package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/udborets/chat-app/server/internal/handlers"
	"github.com/udborets/chat-app/server/internal/repository"
)

func StartServer() {
	// config for database
	config := "host=localhost port=5432 user=postgres password=1234 dbname=chat-app sslmode=disable"
	repository.InitDB(config)

	auth := handlers.NewHTTP()
	ws := handlers.NewWebsHandler()

	router := gin.Default()
	auth.InitRoutes(router)
	ws.InitWebsock(router)

	router.Run(":1773")
}
