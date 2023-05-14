package models

type Message struct {
	MessageId string `json:"message_id"`
	Text      string `json:"text" binding:"required"`
	IsSeen    bool   `json:"is_seen" binding:"required"`
	CreatedAt int    `json:"created_at" binding:"required"`
	UpdatedAt int    `json:"updated_at" binding:"required"`
}

type ChatMessage struct {
	ChatId    string `json:"chat_id"`
	MessageId string `json:"message_id"`
}
