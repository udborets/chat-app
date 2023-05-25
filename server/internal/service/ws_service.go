package service

import (
	"github.com/udborets/chat-app/server/internal/models"
	"github.com/udborets/chat-app/server/internal/repository"
	"net/http"
)

type IWebsBLogic interface {
	ConnectToChats(mapOfRooms *models.RoomsMap, client *models.Client, userId int) (int, string, error)
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
