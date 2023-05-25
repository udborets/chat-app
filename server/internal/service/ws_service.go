package service

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/udborets/chat-app/server/internal/models"
	"github.com/udborets/chat-app/server/internal/repository"
	"log"
	"net/http"
	"time"
)

type IWebsBLogic interface {
	ConnectToChats(mapOfRooms *models.RoomsMap, client *models.Client, userId int) (int, string, error)
	ConnectToChat(mapOfRooms *models.RoomsMap, client *models.Client, chatId, userId int) (int, string, error)
	ReadMessages(mapOfRooms *models.RoomsMap, client *models.Client, chatId int)
	CreateRoom() (int, string, error)
	//GetChats(userId int) (int, string, error)
	//GetRoomsByUserId(userId int) (interface{}, string, error)
	//CreateRoom(users []int) (int, string, error)
}

type WebsBLogic struct {
	websRepository repository.IWebsRepository
}

func NewWebsBLogic() *WebsBLogic {
	return &WebsBLogic{
		websRepository: repository.NewWebsRepository(),
	}
}

func (b *WebsBLogic) ConnectToChats(mapOfRooms *models.RoomsMap, client *models.Client, userId int) (int, string, error) {
	chats, err := b.websRepository.GetChats(userId)
	if err != nil {
		return http.StatusInternalServerError, "error on getting chats from database", err
	}

	for _, chat := range chats.([]int) {
		mapOfRooms.Lock()
		if room, ok := mapOfRooms.Rooms[chat]; ok {
			room.Lock()
			room.Clients[client] = true
			room.Unlock()
		} else {
			newRoom := models.NewRoom(chat)
			newRoom.Lock()
			newRoom.Clients[client] = true
			mapOfRooms.Rooms[chat] = newRoom
			newRoom.Unlock()
		}
		mapOfRooms.Unlock()
	}

	return http.StatusOK, "successfully connected", nil
}

func (b *WebsBLogic) ConnectToChat(mapOfRooms *models.RoomsMap, client *models.Client, userId, chatId int) (int, string, error) {
	err := b.websRepository.CheckChat(chatId)
	if err != nil {
		return http.StatusBadRequest, fmt.Sprintf("no chat with id: %d", chatId), err
	}

	msg, err := b.websRepository.AddUserToChat(userId, chatId)
	if err != nil {
		return http.StatusBadRequest, msg, err
	}

	mapOfRooms.Lock()
	defer mapOfRooms.Unlock()
	if room, ok := mapOfRooms.Rooms[chatId]; ok {
		room.Lock()
		room.Clients[client] = true
		room.Unlock()
	} else {
		newRoom := models.NewRoom(chatId)
		newRoom.Lock()
		newRoom.Clients[client] = true
		mapOfRooms.Rooms[chatId] = newRoom
		newRoom.Unlock()
	}

	return http.StatusOK, "successfully connected", nil
}

func (b *WebsBLogic) ReadMessages(mapOfRooms *models.RoomsMap, client *models.Client, chatId int) {
	room := mapOfRooms.Rooms[chatId]
	defer func() {
		room.RemoveClient(client)
		if len(room.Clients) == 0 {
			mapOfRooms.RemoveRoom(room.RoomId)
		}
	}()

	for {
		_, payload, err := client.Connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message %v: ", err)
			}
			break
		}

		for chatter := range room.Clients {
			chatter.Messages <- payload
		}
	}
}

func (b *WebsBLogic) WriteMessages(mapOfRooms *models.RoomsMap, client *models.Client, chatId int) {
	room := mapOfRooms.Rooms[chatId]
	defer func() {
		room.RemoveClient(client)
		if len(room.Clients) == 0 {
			mapOfRooms.RemoveRoom(room.RoomId)
		}
	}()

	for {
		select {
		case message, ok := <-client.Messages:
			if !ok {
				if err := client.Connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Printf("connection closed: %v", err)
				}
				return
			}
			if err := client.Connection.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Printf("error on sending message: %v", err)
			}
		}
	}
}

func (b *WebsBLogic) CreateRoom() (int, string, error) {
	chat := &models.Chat{LastMessage: nil, UpdatedAt: time.Now().Unix(), CreatedAt: time.Now().Unix()}

	chatId, msg, err := b.websRepository.NewRoom(chat)
	if err != nil {
		return http.StatusBadRequest, msg, err
	}
	return chatId, "", nil
}

//func (b *WebsBLogic) CreateRoom(usersId []int) (int, string, error) {
//	if len(usersId) < 2 {
//		return http.StatusBadRequest, "not enough user id, minimum is 2", errors.New("not enough user id, minimum is 2")
//	}
//
//	msg, err := b.websRepository.CheckUsers(usersId)
//	if err != nil {
//		return http.StatusBadRequest, msg, err
//	}
//
//	chat := models.Chat{
//		LastMessage: nil,
//		CreatedAt:   time.Now().Unix(),
//		UpdatedAt:   time.Now().Unix(),
//	}
//
//	chatId, msg, err := b.websRepository.NewRoom(chat)
//	if err != nil {
//		return 0, msg, err
//	}
//
//	msg, err = b.websRepository.ConnectUsersToChat(usersId, chatId)
//	if err != nil {
//		return 0, msg, err
//	}
//
//	return chatId, "successfully create new chat", nil
//}

//func (b *WebsBLogic) GetChats(userId int) (interface{}, int, string, error) {
//
//}

//func (b *WebsBLogic) GetRoomsByUserId(userId int) (interface{}, string, error) {
//	rooms, err := b.websRepository.GetRooms(userId)
//	if err != nil {
//		return nil, "couldn't get rooms by userId", err
//	}
//	return rooms, "successfully get rooms by userId", err
//}
//
