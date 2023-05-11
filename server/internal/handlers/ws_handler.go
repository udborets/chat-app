package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	websocket2 "github.com/udborets/chat-app/server/internal/service"
)

type WebsHandler struct {
	websBLogic websocket2.IWebsBLogic
}

func NewWebsHandler() *WebsHandler {
	return &WebsHandler{
		websBLogic: websocket2.NewWebsBLogic(),
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *WebsHandler) InitWebsock(router *gin.Engine) {

}
