package models

type Chat struct {
	ChatId      string `json:"chat_id"`
	LastMessage string `json:"last_message"`
	CreatedAt   int64  `json:"created_at" binding:"required"`
	UpdatedAt   int64  `json:"updated_at" binding:"required"`
}

type UserChat struct {
	UserId string `json:"user_id"`
	ChatId string `json:"chat_id"`
}

type Friends struct {
	UserId   string `json:"user_id"`
	FriendId string `json:"friend_id"`
}

type ChatCreateInput struct {
	Users []int `json:"users"`
}
