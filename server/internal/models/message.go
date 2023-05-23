package models

type Message struct {
	MessageId int    `json:"message_id"`
	Text      string `json:"text" binding:"required"`
	IsSeen    bool   `json:"is_seen" binding:"required"`
	CreatedAt int64  `json:"created_at" binding:"required"`
	UpdatedAt int64  `json:"updated_at" binding:"required"`
}

type MessageValidateOutput struct {
	MessageId int    `json:"message_id"`
	Text      string `json:"text"`
	IsSeen    bool   `json:"is_seen"`
}

type ChatMessage struct {
	ChatId    string `json:"chat_id"`
	MessageId string `json:"message_id"`
}
