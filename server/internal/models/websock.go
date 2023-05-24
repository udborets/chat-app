package models

import (
	"github.com/gorilla/websocket"
	"sync"
)

var RoomsMap struct {
	Rooms map[int]*Room
	sync.Mutex
}

type Room struct {
	Clients map[*Client]bool
	RoomId  int
	sync.Mutex
}

func NewRoom(client *Client, roomId int) *Room {
	mp := make(map[*Client]bool)
	mp[client] = true
	room := &Room{Clients: mp, RoomId: roomId}
	RoomsMap.Rooms[roomId] = room
	return room
}

func AddClient(client *Client, room *Room) {
	client.Rooms = append(client.Rooms, room)
	room.Clients[client] = true
}

type Client struct {
	Connection *websocket.Conn
	Rooms      []*Room
	Messages   chan []byte
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{Connection: conn, Rooms: make([]*Room, 0), Messages: make(chan []byte)}
}

func DeleteClient(client *Client) {
	for _, room := range client.Rooms {
		room.Lock()
		delete(room.Clients, client)

		if len(room.Clients) == 0 {
			RoomsMap.Lock()
			delete(RoomsMap.Rooms, room.RoomId)
			RoomsMap.Unlock()
		}

		room.Unlock()
	}
	client.Connection.Close()
}
