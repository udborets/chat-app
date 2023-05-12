package repository

import (
	"database/sql"
	"fmt"
	"github.com/udborets/chat-app/server/internal/models"
)

type IWebsRepository interface {
	GetRooms(userId int) ([]models.Chat, error)
	NewRoom(chat models.Chat) (int, string, error)
	ConnectUsersToChat(users []int, chatId int) (string, error)
}

type WebsRepository struct {
	db *sql.DB
}

func NewWebsRepository() *WebsRepository {
	return &WebsRepository{
		db: database,
	}
}

func (r *WebsRepository) GetRooms(userId int) ([]models.Chat, error) {
	rows, err := r.db.Query("SELECT chat_id, last_message, created_at, updated_at"+
		"FROM user_chat INNER JOIN chat USING(chat_id)"+
		"WHERE user_id = $1 ORDER BY updated_at DESC", userId)
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

func (r *WebsRepository) NewRoom(chat models.Chat) (int, string, error) {
	var chatId int

	err := r.db.QueryRow("insert into \"chat\" (last_message, created_at, updated_at)"+
		"values ($1,$2,$3) returning chat_id", chat.LastMessage, chat.CreatedAt, chat.UpdatedAt).Scan(&chatId)
	if err != nil {
		return 0, "couldn't add chat to 'chat' database", err
	}

	return chatId, "successfully added chat", nil
}

func (r *WebsRepository) ConnectUsersToChat(users []int, chatId int) (string, error) {
	for _, user := range users {
		_, err := r.db.Exec("insert into \"user_chat\" (user_id, chat_id)"+
			"values ($1, $2)", user, chatId)
		if err != nil {
			return fmt.Sprintf("couldn't add user with id: %s", user), err
		}
	}
	return fmt.Sprintf("all users connected to chat with chat_id: %s", chatId), nil
}
