package service

import (
	"github.com/udborets/chat-app/server/internal/repository"
)

type IWebsBLogic interface {
	GetRoomsByUserId(userId int) (interface{}, string, error)
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
	//roomsId, err := b.websRepository.GetRoomsId(userId)
	//if err != nil {
	//	return "incorrect userId", err
	//}
	rooms, err := b.websRepository.GetRooms(userId)
	if err != nil {
		return nil, "couldn't get rooms by userId", err
	}
	return rooms, "successfully get rooms by userId", err
}
