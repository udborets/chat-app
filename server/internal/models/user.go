package models

type UserSignUpInput struct {
	Name      string      `json:"name" binding:"required"`
	Email     interface{} `json:"email"`
	Phone     interface{} `json:"phone"`
	Password  string      `json:"password" binding:"required"`
	AvatarURL interface{} `json:"avatarURL"`
}

type User struct {
	UserId    int         `json:"id"`
	Name      string      `json:"name" binding:"required"`
	HashPass  string      `json:"hash_password" binding:"required"`
	Email     interface{} `json:"email"`
	Phone     interface{} `json:"phone"`
	AvatarURL interface{} `json:"avatar_url"`
	LastSeen  interface{} `json:"last_seen"`
	CreatedAt int64       `json:"created_at" binding:"required"`
	UpdatedAt int64       `json:"updated_at" binding:"required"`
}

type UserSignInInput struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password" binding:"required"`
}
