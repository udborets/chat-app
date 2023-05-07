package repository

import (
	"database/sql"
	"github.com/udborets/chat-app/server/internal/utilities"
	"log"
)

type IAuthDB interface {
	//GenerateJWTToken(email, password string) (string, error)
	//CheckPass(email, password string) bool
	AddUser(user utilities.User) (error, string)
}

type AuthDB struct {
	db *sql.DB
}

func NewAuthDB(config string) *AuthDB {
	db, err := sql.Open("postgres", config)
	if err != nil {
		log.Printf("couldn't connect to postgres with config: %s", config)
	}
	return &AuthDB{db: db}
}

//func (a *AuthDB) GenerateJWTToken(email, password string) (string, error) {
//
//}
//
//func (a *AuthDB) CheckPass(email, password string) (error, string) {
//	row := a.db.QueryRow(`select * from "auth" where email=$1`, email)
//}

func (a *AuthDB) AddUser(user utilities.User) (error, string) {
	_, err := a.db.Exec(`insert into "auth" (name, email, phone, password, avatarURL) values ($1,$2,$3,$4,$5)`,
		user.Name, user.Email, user.Phone, user.HashPass, user.AvatarURL)
	if err != nil {
		return err, "error on adding user to 'authorization' database"
	}
	_, err = a.db.Exec(`insert into "User" (name, hashPassword, email, phone, avatarURL, createdAt) values ($1,$2,$3,$4,$5,$6)`,
		user.Name, user.HashPass, user.Email, user.Phone, user.AvatarURL, user.CreatedAt)
	if err != nil {
		return err, "error on adding user to 'user' database"
	}
	return nil, ""
}
