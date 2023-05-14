package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/udborets/chat-app/server/internal/models"
	"github.com/udborets/chat-app/server/internal/responses"
	"github.com/udborets/chat-app/server/internal/service"
	"net/http"
	"os"
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

	websock.GET("/chats/:userId", h.getRooms)
	websock.GET("/newChat", h.newRoom)
	websock.GET("/joinChat/:chatId", h.joinRoom)
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

func (h *WebsHandler) newRoom(ctx *gin.Context) {
	var inp models.ChatCreateInput

	if err := ctx.BindJSON(&inp); err != nil {
		responses.NewResponse(ctx, http.StatusBadRequest, "newRoom receive array of user's_id", err)
		return
	}

	chatId, msg, err := h.websBLogic.CreateRoom(inp.Users)
	if err != nil {
		responses.NewResponse(ctx, http.StatusBadRequest, msg, err)
		return
	}

	responses.NewResponse(ctx, http.StatusOK, msg, err)

	redirectUrl := fmt.Sprintf("http://localhost:%s/ws/joinRoom/%d", os.Getenv("PORT"), chatId)
	ctx.Redirect(http.StatusFound, redirectUrl)
}

func (h *WebsHandler) joinRoom(ctx *gin.Context) {

}
