package models

type Message struct {
	MessageId string `json:"id"`
	Text      string `json:"text" binding:"required"`
	IsSeen    bool   `json:"isSeen" binding:"required"`
	CreatedAt int    `json:"createdAt" binding:"required"`
	UpdatedAt int    `json:"updatedAt" binding:"required"`
}

type ChatMessage struct {
	ChatId    string `json:"chat_id"`
	MessageId string `json:"message_id"`
}
