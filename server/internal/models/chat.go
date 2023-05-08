package models

type Chat struct {
	ChatId    string `json:"id"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
}

type UserChat struct {
	UserId string `json:"user_id"`
	ChatId string `json:"chat_id"`
}

type Friends struct {
	UserId   string `json:"user_id"`
	FriendId string `json:"friend_id"`
}
