package models

type Chat struct {
	ChatId      string `json:"id"`
	LastMessage string `json:"lastMessage"`
	CreatedAt   int    `json:"createdAt" binding:"required"`
	UpdatedAt   int    `json:"updatedAt" binding:"required"`
}

type UserChat struct {
	UserId string `json:"user_id"`
	ChatId string `json:"chat_id"`
}

type Friends struct {
	UserId   string `json:"user_id"`
	FriendId string `json:"friend_id"`
}
