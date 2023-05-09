package repository

import (
	"database/sql"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
	"github.com/udborets/chat-app/server/internal/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

type IAuthDB interface {
	AddUser(user models.User) (error, string)
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

func (a *AuthDB) CheckPassByEmail(email, pass string) (error, string) {
	row := a.db.QueryRow(`select hash_password, user_id from "auth" where email = $1`, email)

	var corrPass string
	var id int

	fmt.Println(row)
	err := row.Scan(&corrPass, &id)
	if err != nil {
		return err, "no user with this email"
	}

	if err = bcrypt.CompareHashAndPassword([]byte(corrPass), []byte(pass)); err != nil {
		return err, "incorrect password"
	}

	err, token := createJWTToken(id)
	if err != nil {
		return err, err.Error()
	}

	return nil, token
}

func (a *AuthDB) CheckPassByName(name, pass string) (error, string) {
	row := a.db.QueryRow(`select password, userId from "auth" where name = $1`, name)

	var corrPass string
	var id int

	err := row.Scan(&corrPass, &id)
	if err != nil {
		return err, "no user with this name"
	}

	if err = bcrypt.CompareHashAndPassword([]byte(corrPass), []byte(pass)); err != nil {
		return err, "incorrect password"
	}

	err, token := createJWTToken(id)
	if err != nil {
		return err, err.Error()
	}

	return nil, token
}

func (a *AuthDB) CheckPassByPhone(phone, pass string) (error, string) {
	row := a.db.QueryRow(`select password, userId from "auth" where phone = $1`, phone)

	var corrPass string
	var id int
	err := row.Scan(&corrPass, &id)
	if err != nil {
		return err, "no user with this phone"
	}

	if err = bcrypt.CompareHashAndPassword([]byte(corrPass), []byte(pass)); err != nil {
		return err, "incorrect password"
	}

	err, token := createJWTToken(id)
	if err != nil {
		return err, err.Error()
	}

	return nil, token
}

func createJWTToken(userId int) (error, string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return err, ""
	}

	return nil, tokenString
}
