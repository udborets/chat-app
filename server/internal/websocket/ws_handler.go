package websocket

import "github.com/gorilla/websocket"

type WsHandler struct {
	wsBLogic IWsBLogic
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
