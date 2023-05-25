package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/udborets/chat-app/server/internal/handlers"
	"github.com/udborets/chat-app/server/internal/repository"
	"os"
)

func StartServer() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, pass, name, sslmode)

	repository.InitDB(config)

	auth := handlers.NewHTTP()
	ws := handlers.NewWebsHandler()

	router := gin.Default()
	auth.InitRoutes(router)
	ws.InitWebsock(router)

	router.Run(":" + os.Getenv("PORT"))
}
