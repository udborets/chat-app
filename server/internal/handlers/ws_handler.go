package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/udborets/chat-app/server/internal/responses"
	"github.com/udborets/chat-app/server/internal/service"
	"net/http"
	"strconv"
	"sync"
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

	client := NewClient(conn)
	statusCode, msg, err := h.websBLogic.ConnectToChats(client, userId)
	if err != nil {
		responses.NewResponse(ctx, statusCode, msg, err)
		return
	}

	client.ReadMessage()
	//go client.WriteMessage()
}

func (c *Client) ReadMessage() {
	defer func() {
		c.DeleteClient()
	}()

	for {
		_, payload, err := c.Connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("error reading message: %v", err)
			}
			break
		}

		for
	}
}

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

var RoomsMap struct {
	Rooms map[int]*Room
	sync.Mutex
}

type Room struct {
	Clients map[*Client]bool
	RoomId  int
	sync.Mutex
}

func NewRoom(client *Client, roomId int) *Room {
	mp := make(map[*Client]bool)
	mp[client] = true
	room := &Room{Clients: mp, RoomId: roomId}
	RoomsMap.Rooms[roomId] = room
	return room
}

func DeleteRoom(roomId int) {
	RoomsMap.Lock()
	defer RoomsMap.Unlock()
	delete(RoomsMap.Rooms, roomId)
}

func (c *Client) AddClient(room *Room) {
	c.Rooms = append(c.Rooms, room)
	room.Clients[c] = true
}

type Client struct {
	Connection *websocket.Conn
	Rooms      []*Room
	Messages   chan []byte
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{Connection: conn, Rooms: make([]*Room, 0), Messages: make(chan []byte)}
}

func (c *Client) DeleteClient() {
	for _, room := range c.Rooms {
		room.Lock()
		delete(room.Clients, c)

		if len(room.Clients) == 0 {
			RoomsMap.Lock()
			delete(RoomsMap.Rooms, room.RoomId)
			RoomsMap.Unlock()
		}

		room.Unlock()
	}
	c.Connection.Close()
}
