package models

type Chat struct {
	ChatId      string      `json:"chat_id"`
	LastMessage interface{} `json:"last_message"`
	CreatedAt   int64       `json:"created_at" binding:"required"`
	UpdatedAt   int64       `json:"updated_at" binding:"required"`
}

//type ChatValidateOutput struct {
//	ChatId      string      `json:"chat_id"`
//	LastMessage interface{} `json:"last_message"`
//}

type ChatValidateOutput struct {
	ChatId   int                     `json:"chat_id"`
	Messages []MessageValidateOutput `json:"messages"`
}

type UserChat struct {
	UserId int `json:"user_id"`
	ChatId int `json:"chat_id"`
}

type Friends struct {
	UserId   int `json:"user_id"`
	FriendId int `json:"friend_id"`
}

type ChatCreateInput struct {
	Users []int `json:"users"`
}
