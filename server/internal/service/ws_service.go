package service

import (
	"github.com/udborets/chat-app/server/internal/models"
	"github.com/udborets/chat-app/server/internal/repository"
	"time"
)

type IWebsBLogic interface {
	GetRoomsByUserId(userId int) (interface{}, string, error)
	CreateRoom(users []int) (int, string, error)
}

type WebsBLogic struct {
	websRepository repository.IWebsRepository
}

func NewWebsBLogic() *WebsBLogic {
	return &WebsBLogic{
		websRepository: repository.NewWebsRepository(),
	}
}

func (b *WebsBLogic) GetRoomsByUserId(userId int) (interface{}, string, error) {
	rooms, err := b.websRepository.GetRooms(userId)
	if err != nil {
		return nil, "couldn't get rooms by userId", err
	}
	return rooms, "successfully get rooms by userId", err
}

func (b *WebsBLogic) CreateRoom(usersId []int) (int, string, error) {
	chat := models.Chat{
		LastMessage: "",
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}

	chatId, msg, err := b.websRepository.NewRoom(chat)
	if err != nil {
		return 0, msg, err
	}

	msg, err = b.websRepository.ConnectUsersToChat(usersId, chatId)
	if err != nil {
		return 0, msg, err
	}

	return chatId, "successful", nil
}
