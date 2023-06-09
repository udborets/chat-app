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

//go:generate mockgen -source=websock.go -destination=mocks/websock_mock.go

type IWebsService interface {
	ConnectToChats(mapOfRooms *models.RoomsMap, client *models.Client, userId int) (int, string, error)
	CheckParams(userId, chatId int) (string, error)
	ConnectToChat(mapOfRooms *models.RoomsMap, client *models.Client, chatId, userId int) (int, string, error)
	ReadMessages(mapOfRooms *models.RoomsMap, client *models.Client, chatId, userId int)
	WriteMessages(mapOfRooms *models.RoomsMap, client *models.Client, chatId int)
	CreateRoom() (int, string, error)
}

type WebsService struct {
	websRepository repository.IWebsRepository
}

func NewWebsService() *WebsService {
	return &WebsService{
		websRepository: repository.NewWebsRepository(),
	}
}

func (b *WebsService) ConnectToChats(mapOfRooms *models.RoomsMap, client *models.Client, userId int) (int, string, error) {
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

func (b *WebsService) ConnectToChat(mapOfRooms *models.RoomsMap, client *models.Client, userId, chatId int) (int, string, error) {
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

func (b *WebsService) CheckParams(userId, chatId int) (string, error) {
	err := b.websRepository.CheckUser(userId)
	if err != nil {
		return fmt.Sprintf("no user with id: %d", userId), err
	}

	err = b.websRepository.CheckChat(chatId)
	if err != nil {
		return fmt.Sprintf("no chat with id: %d", chatId), err
	}
	return "", nil
}

func (b *WebsService) ReadMessages(mapOfRooms *models.RoomsMap, client *models.Client, chatId, userId int) {
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
			fmt.Println(err)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message %v: ", err)
			}
			break
		}

		newMessage := models.Message{
			ChatId:    chatId,
			Text:      string(payload),
			SenderId:  userId,
			IsSeen:    false,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}

		err = b.websRepository.AddMessage(&newMessage)
		if err != nil {
			log.Fatal(err)
			break
		}

		//fmt.Printf("sending message: %s to clients: %v\n", string(payload), room.Clients)
		for chatter := range room.Clients {
			chatter.Messages <- newMessage
		}
	}
}

func (b *WebsService) WriteMessages(mapOfRooms *models.RoomsMap, client *models.Client, chatId int) {
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
			if err := client.Connection.WriteMessage(websocket.TextMessage, []byte(message.Text)); err != nil {
				log.Printf("error on sending message: %v", err)
			}

			err := b.websRepository.ChangeMessageSeen(message.MessageId)
			if err != nil {
				log.Printf("error on changing is_seen value of message")
			}
		}
	}
}

func (b *WebsService) CreateRoom() (int, string, error) {
	chat := &models.Chat{LastMessageId: nil, UpdatedAt: time.Now().Unix(), CreatedAt: time.Now().Unix()}

	chatId, msg, err := b.websRepository.NewRoom(chat)
	if err != nil {
		return http.StatusBadRequest, msg, err
	}
	return chatId, "", nil
}

//func (b *WebsService) CreateRoom(usersId []int) (int, string, error) {
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

//func (b *WebsService) GetChats(userId int) (interface{}, int, string, error) {
//
//}

//func (b *WebsService) GetRoomsByUserId(userId int) (interface{}, string, error) {
//	rooms, err := b.websRepository.GetRooms(userId)
//	if err != nil {
//		return nil, "couldn't get rooms by userId", err
//	}
//	return rooms, "successfully get rooms by userId", err
//}
//
