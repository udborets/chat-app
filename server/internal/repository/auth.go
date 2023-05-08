package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/udborets/chat-app/server/internal/models"
	"log"
)

type IAuthDB interface {
	AddUser(user models.User) (error, string)
	//CheckPassByName(name, pass string) (error, string)
	//CheckPassByEmail(email, pass string) (error, string)
	//CheckPassByPhone(email, pass string) (error, string)
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

func (a *AuthDB) AddUser(user models.User) (error, string) {
	var addedId int

	err := a.db.QueryRow(`insert into "user" (name, hash_password, email, phone, avatar_url, created_at) values ($1,$2,$3,$4,$5,$6) returning user_id`,
		user.Name, user.HashPass, user.Email, user.Phone, user.AvatarURL, user.CreatedAt).Scan(&addedId)
	if err != nil {
		return err, "error taking id from added user"
	}

	_, err = a.db.Exec(`insert into "auth" (name, email, phone, hash_password, user_id) values ($1,$2,$3,$4,$5)`,
		user.Name, user.Email, user.Phone, user.HashPass, addedId)
	if err != nil {
		return err, "error on adding user to 'authorization' database"
	}
	return nil, ""
}

//func (a *AuthDB) CheckPassByEmail(email, pass string) (error, string) {
//	row := a.db.QueryRow(`select password, user_id from "auth" where email = $1`, email)
//
//	var corrPass string
//	var uid string
//
//	err := row.Scan(&corrPass, &uid)
//	if err != nil {
//		return err, "no user with this email"
//	}
//	bcrypt.CompareHashAndPassword(corrPass, pass)
//	if corrPass != pass {
//		return errors.New("incorrect password"), "password is incorrect"
//	}
//	return nil, ""
//}
//
//func (a *AuthDB) CheckPassByName(name, pass string) (error, string) {
//	row := a.db.QueryRow(`select password, userId from "auth" where name = $1`, name)
//
//	var corrPass string
//	var uid string
//	err := row.Scan(&corrPass, &uid)
//	if err != nil {
//		return err, "no user with this name"
//	}
//
//	if corrPass != pass {
//		return errors.New("incorrect password"), "password is incorrect"
//	}
//	return nil, ""
//}
//
//func (a *AuthDB) CheckPassByPhone(phone, pass string) (error, string) {
//	row := a.db.QueryRow(`select password, userId from "auth" where phone = $1`, phone)
//
//	var corrPass string
//	var uid string
//	err := row.Scan(&corrPass, &uid)
//	if err != nil {
//		return err, "no user with this phone"
//	}
//
//	if corrPass != pass {
//		return errors.New("incorrect password"), "password is incorrect"
//	}
//	return nil, ""
//}
