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

var mapOfRooms *models.RoomsMap

func (h *WebsHandler) InitWebsock(router *gin.Engine) {
	mapOfRooms = models.NewRoomsMap()

	router.GET("/ws:userId", h.connect)

	websock := router.Group("/ws")

	//websock.GET("/chats/:userId", h.getRooms)
	websock.GET("/newChat", h.newRoom)
	websock.GET("/chat", h.joinRoom) // ws://localhost/ws/chat?userId=4&chatId=3
}

func (h *WebsHandler) connect(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		responses.NewResponse(ctx, http.StatusBadRequest, ":userId require int", err)
		return
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		responses.NewResponse(ctx, http.StatusBadRequest, "ws connection required", err)
		return
	}
	defer conn.Close()

	client := models.NewClient(conn)
	statusCode, msg, err := h.websBLogic.ConnectToChats(mapOfRooms, client, userId)
	if err != nil {
		responses.NewResponse(ctx, statusCode, msg, err)
		return
	}
}

func (h *WebsHandler) newRoom(ctx *gin.Context) {
	chatId, msg, err := h.websBLogic.CreateRoom()
	if err != nil {
		responses.NewResponse(ctx, http.StatusBadRequest, msg, err)
		return
	}
	ctx.JSON(http.StatusOK, chatId)
}

func (h *WebsHandler) joinRoom(ctx *gin.Context) { // ws://localhost:8080/ws/chat?userId=5&chatId=3
	userId, err := strconv.Atoi(ctx.Query("userId"))
	if err != nil {
		responses.NewResponse(ctx, http.StatusBadRequest, "user require integer", err)
		return
	}

	chatId, err := strconv.Atoi(ctx.Query("chatId"))
	if err != nil {
		responses.NewResponse(ctx, http.StatusBadRequest, "chatId require integer", err)
		return
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		responses.NewResponse(ctx, http.StatusBadRequest, "websocket connection required", err)
		return
	}
	//defer conn.Close()

	client := models.NewClient(conn)
	statusCode, msg, err := h.websBLogic.ConnectToChat(mapOfRooms, client, userId, chatId)
	if err != nil {
		responses.NewResponse(ctx, statusCode, msg, err)
		return
	}

	go h.websBLogic.ReadMessages(mapOfRooms, client, chatId, userId)
	go h.websBLogic.WriteMessages(mapOfRooms, client, chatId)
}

//func (h *WebsHandler) newRoom(ctx *gin.Context) {
//	var inp models.ChatCreateInput
//
//	if err := ctx.BindJSON(&inp); err != nil {
//		responses.NewResponse(ctx, http.StatusBadRequest, "newRoom receive array of user's_id", err)
//		return
//	}
//
//	chatId, msg, err := h.websBLogic.CreateRoom(inp.Users)
//	if err != nil {
//		responses.NewResponse(ctx, http.StatusBadRequest, msg, err)
//		return
//	}
//
//	responses.NewResponse(ctx, http.StatusOK, msg, err)
//
//	redirectUrl := fmt.Sprintf("http://localhost:%s/ws/joinRoom/%d", os.Getenv("PORT"), chatId)
//	ctx.Redirect(http.StatusFound, redirectUrl)
//}

//func (h *WebsHandler) getRooms(ctx *gin.Context) {
//	userId, err := strconv.Atoi(ctx.Param("userId"))
//	if err != nil {
//		responses.NewResponse(ctx, http.StatusBadRequest, ":userId require int", err)
//		return
//	}
//
//	rooms, msg, err := h.websBLogic.GetRoomsByUserId(userId)
//	if err != nil {
//		responses.NewResponse(ctx, http.StatusInternalServerError, msg, err)
//	}
//
//	ctx.Set("rooms", rooms.([]models.Chat))
//}
//
