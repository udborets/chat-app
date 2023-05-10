package websocket

type IWsBLogic interface{}

type WsBLogic struct {
	wsRepository IWsRepository
}
