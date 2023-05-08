package repository

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/udborets/chat-app/server/internal/utilities"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type IAuthDB interface {
	AddUser(user utilities.User) (error, string)
	CheckPassByName(name, pass string) (error, string)
	CheckPassByEmail(email, pass string) (error, string)
	CheckPassByPhone(email, pass string) (error, string)
}

type AuthDB struct {
	db *sql.DB
}

func NewAuthDB(config string) *AuthDB {
	db, err := sql.Open("postgres", config)
	if err != nil {
		fmt.Println(err.Error())
		log.Printf("couldn't connect to postgres with config: %s", config)
	}
	return &AuthDB{db: db}
}

func (a *AuthDB) AddUser(user utilities.User) (error, string) {
	_, err := a.db.Exec(`insert into "User" (name, hashPassword, email, phone, avatarURL, createdAt) values ($1,$2,$3,$4,$5,$6)`,
		user.Name, user.HashPass, user.Email, user.Phone, user.AvatarURL, user.CreatedAt)
	if err != nil {
		return err, "error on adding user to 'user' database"
	}
	_, err = a.db.Exec(`insert into "auth" (name, email, phone, password) values ($1,$2,$3,$4)`,
		user.Name, user.Email, user.Phone, user.HashPass)
	if err != nil {
		return err, "error on adding user to 'authorization' database"
	}
	return nil, ""
}

func (a *AuthDB) CheckPassByEmail(email, pass string) (error, string) {
	row := a.db.QueryRow(`select password, userId from "auth" where email = $1`, email)

	var corrPass string
	var uid string
	err := row.Scan(&corrPass, &uid)
	if err != nil {
		return err, "no user with this email"
	}
	bcrypt.CompareHashAndPassword(corrPass, pass)
	if corrPass != pass {
		return errors.New("incorrect password"), "password is incorrect"
	}
	return nil, ""
}

func (a *AuthDB) CheckPassByName(name, pass string) (error, string) {
	row := a.db.QueryRow(`select password, userId from "auth" where name = $1`, name)

	var corrPass string
	var uid string
	err := row.Scan(&corrPass, &uid)
	if err != nil {
		return err, "no user with this name"
	}

	if corrPass != pass {
		return errors.New("incorrect password"), "password is incorrect"
	}
	return nil, ""
}

func (a *AuthDB) CheckPassByPhone(phone, pass string) (error, string) {
	row := a.db.QueryRow(`select password, userId from "auth" where phone = $1`, phone)

	var corrPass string
	var uid string
	err := row.Scan(&corrPass, &uid)
	if err != nil {
		return err, "no user with this phone"
	}

	if corrPass != pass {
		return errors.New("incorrect password"), "password is incorrect"
	}
	return nil, ""
}
