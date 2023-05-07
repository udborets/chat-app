package utilities

type Chat struct {
	Id            string    `json:"id"`
	CompanionUser User      `json:"companionUser"`
	Messages      []Message `json:"messages"`
	CreatedAt     int       `json:"createdAt"`
	UpdatedAt     int       `json:"updatedAt"`
}
