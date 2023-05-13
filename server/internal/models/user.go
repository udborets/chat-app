package models

type UserSignUpInput struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password" binding:"required"`
	AvatarURL string `json:"avatarURL"`
}

type User struct {
	UserId    string `json:"id"`
	Name      string `json:"name" binding:"required"`
	HashPass  string `json:"password" binding:"required"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	AvatarURL string `json:"avatarURL"`
	CreatedAt int64  `json:"createdAt" binding:"required"`
	UpdatedAt int64  `json:"updatedAt" binding:"required"`
	LastSeen  int64  `json:"lastSeen" binding:"required"`
}

type UserSignInInput struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password" binding:"required"`
}
