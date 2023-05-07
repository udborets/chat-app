package utilities

type Message struct {
	Id        string `json:"id"`
	Text      string `json:"text"`
	SenderId  string `json:"senderId"`
	IsSeen    bool   `json:"isSeen"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
}
