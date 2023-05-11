package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebsHandler struct {
	websBLogic IWebsBLogic
}

func NewWebsHandler() *WebsHandler {
	return &WebsHandler{
		websBLogic: NewWebsBLogic(),
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *WebsHandler) InitWebsock(router *gin.Engine) {

}
