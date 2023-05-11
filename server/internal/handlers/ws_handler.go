package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/udborets/chat-app/server/internal/models"
	"github.com/udborets/chat-app/server/internal/responses"
	"github.com/udborets/chat-app/server/internal/service"
	"net/http"
	"strconv"
)

type WebsHandler struct {
	websBLogic service.IWebsBLogic
}

func NewWebsHandler() *WebsHandler {
	return &WebsHandler{
		websBLogic: service.NewWebsBLogic(),
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (h *WebsHandler) InitWebsock(router *gin.Engine) {
	websock := router.Group("/ws")

	websock.GET("/rooms/:userId", h.getRooms)
	websock.GET("/joinRoom/:roomId", h.joinRoom)
}

func (h *WebsHandler) getRooms(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		responses.NewResponse(ctx, http.StatusBadRequest, ":userId require int", err)
		return
	}

	rooms, msg, err := h.websBLogic.GetRoomsByUserId(userId)
	if err != nil {
		responses.NewResponse(ctx, http.StatusInternalServerError, msg, err)
	}

	ctx.Set("rooms", rooms.([]models.Chat))
}

func (h *WebsHandler) joinRoom(ctx *gin.Context) {

}
