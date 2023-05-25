package models

import (
	"github.com/gorilla/websocket"
	"sync"
)

type RoomsMap struct {
	Rooms map[int]*Room
	sync.Mutex
}

type Room struct {
	RoomId  int
	Clients map[*Client]bool
	sync.Mutex
}

type Client struct {
	Connection *websocket.Conn
	Rooms      []*Room
	Messages   chan []byte
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{Connection: conn,
		Rooms:    make([]*Room, 0),
		Messages: make(chan []byte),
	}
}

func NewRoom(roomId int) *Room {
	return &Room{RoomId: roomId, Clients: make(map[*Client]bool)}
}

func NewRoomsMap() *RoomsMap {
	return &RoomsMap{Rooms: make(map[int]*Room)}
}

//func NewRoom(client *Client, roomId int) *Room {
//	mp := make(map[*Client]bool)
//	mp[client] = true
//	room := &Room{Clients: mp, RoomId: roomId}
//	return room
//}
//
//func DeleteRoom(roomId int) {
//	RoomsMap.Lock()
//	defer RoomsMap.Unlock()
//	delete(RoomsMap.Rooms, roomId)
//}
//
//func (c *Client) AddClient(room *Room) {
//	c.Rooms = append(c.Rooms, room)
//	room.Clients[c] = true
//}
//
//func NewClient(conn *websocket.Conn) *Client {
//	return &Client{Connection: conn, Rooms: make([]*Room, 0), Messages: make(chan []byte)}
//}
//
//func (c *Client) DeleteClient() {
//	for _, room := range c.Rooms {
//		room.Lock()
//		delete(room.Clients, c)
//
//		if len(room.Clients) == 0 {
//			RoomsMap.Lock()
//			delete(RoomsMap.Rooms, room.RoomId)
//			RoomsMap.Unlock()
//		}
//
//		room.Unlock()
//	}
//	c.Connection.Close()
//}
