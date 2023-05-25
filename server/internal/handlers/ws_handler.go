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

	//websock := router.Group("/ws")
	//
	//websock.GET("/chats/:userId", h.getRooms)
	//websock.GET("/newChat", h.newRoom)
	//websock.GET("/joinChat/:chatId", h.joinRoom)
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

//func (c *Client) ReadMessage() {
//	defer func() {
//		c.DeleteClient()
//	}()
//
//	for {
//		_, _, err := c.Connection.ReadMessage()
//
//		if err != nil {
//			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
//				fmt.Printf("error reading message: %v", err)
//			}
//			break
//		}
//	}
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
//
//func (h *WebsHandler) joinRoom(ctx *gin.Context) {
//
//}