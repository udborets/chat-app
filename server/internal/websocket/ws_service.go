package websocket

type IWebsBLogic interface{}

type WebsBLogic struct {
	websRepository IWebsRepository
}

func NewWebsBLogic() *WebsBLogic {
	return &WebsBLogic{
		websRepository: NewWebsRepository(),
	}
}
