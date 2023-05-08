package models

type UserSignUpInput struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Phone     string `json:"phone"`
	Password  string `json:"password" binding:"required"`
	AvatarURL string `json:"avatarURL"`
}

type User struct {
	UserId    string `json:"id"`
	Name      string `json:"name" binding:"required"`
	HashPass  string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Phone     string `json:"phone"`
	AvatarURL string `json:"avatarURL"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	LastSeen  int64  `json:"lastSeen"`
}

type UserSignInInput struct {
	Login    string `json:"login"`
	Password string `json:"password" binding:"required"`
}

type UserAuth struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCheckEmail struct {
}
