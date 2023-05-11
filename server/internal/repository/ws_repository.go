package repository

import (
	"database/sql"
	"github.com/udborets/chat-app/server/internal/models"
)

type IWebsRepository interface {
	GetRoomsId(userId int) ([]int, error)
	GetRooms(userId int) ([]models.Chat, error)
}

type WebsRepository struct {
	db *sql.DB
}

func NewWebsRepository() *WebsRepository {
	return &WebsRepository{
		db: database,
	}
}

func (r *WebsRepository) GetRoomsId(userId int) ([]int, error) {
	rows, err := r.db.Query("select chat_id from \"user_chat\" where user_id=$1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roomsId := make([]int, 0)

	for rows.Next() {
		var roomId int
		err := rows.Scan(&roomId)
		if err != nil {
			return nil, err
		}
		roomsId = append(roomsId, roomId)
	}

	return roomsId, nil
}

func (r *WebsRepository) GetRooms(userId int) ([]models.Chat, error) {
	rows, err := r.db.Query("SELECT chat_id, last_message, created_at, updated_at"+
		"FROM user_chat INNER JOIN chat USING(chat_id) ORDER BY updated_at DESC"+
		"WHERE user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rooms := make([]models.Chat, 0)

	for rows.Next() {
		var room models.Chat
		err := rows.Scan(&room.ChatId, &room.LastMessage, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}
