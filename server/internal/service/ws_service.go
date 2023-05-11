package service

import "github.com/udborets/chat-app/server/internal/repository"

type IWebsBLogic interface{}

type WebsBLogic struct {
	websRepository repository.IWebsRepository
}

func NewWebsBLogic() *WebsBLogic {
	return &WebsBLogic{
		websRepository: repository.NewWebsRepository(),
	}
}
