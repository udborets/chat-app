package models

type Message struct {
	MessageId string `json:"id"`
	Text      string `json:"text"`
	IsSeen    bool   `json:"isSeen"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
}

type ChatMessage struct {
	ChatId    string `json:"chat_id"`
	MessageId string `json:"message_id"`
}
