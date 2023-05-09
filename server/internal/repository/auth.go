package repository

import (
	"database/sql"
	"errors"
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
	AddUser(user models.User) (string, error)
	CheckPassByName(name, pass string) (string, error)
	CheckPassByEmail(email, pass string) (string, error)
	CheckPassByPhone(email, pass string) (string, error)
	GetUserByID(userId int) (models.User, error)
	CheckUniqUser(user models.UserSignUpInput) (string, error)
}

type AuthDB struct {
	db *sql.DB
}

func NewAuthDB(config string) *AuthDB {
	db, err := sql.Open("postgres", config)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return &AuthDB{db: db}
}

func (a *AuthDB) AddUser(user models.User) (string, error) {
	var addedId int

	err := a.db.QueryRow("insert into \"user\" (name, hash_password, email, phone, avatar_url, created_at, updated_at, last_seen) "+
		"values ($1,$2,$3,$4,$5,$6,$7,$8) returning user_id", user.Name, user.HashPass, user.Email, user.Phone, user.AvatarURL,
		user.CreatedAt, user.UpdatedAt, user.LastSeen).Scan(&addedId)
	if err != nil {
		return "server couldn't add data to database", err
	}

	_, err = a.db.Exec(`insert into "auth" (name, email, phone, hash_password, user_id) values ($1,$2,$3,$4,$5)`,
		user.Name, user.Email, user.Phone, user.HashPass, addedId)
	if err != nil {
		return "server couldn't add data to database", err
	}
	return "", nil
}

func (a *AuthDB) CheckUniqUser(user models.UserSignUpInput) (string, error) {
	if msg := a.checkHelper("email", user.Email); msg != "" {
		return msg, errors.New("server couldn't add row to column with unique columns")
	} else if msg := a.checkHelper("name", user.Name); msg != "" {
		return msg, errors.New("server couldn't add row to column with unique columns")
	} else if msg := a.checkHelper("phone", user.Phone); msg != "" {
		return msg, errors.New("server couldn't add row to column with unique columns")
	}
	return "", nil
}

func (a *AuthDB) checkHelper(column, variable string) string {
	var has bool

	query := fmt.Sprintf("select exists (select * from \"auth\" where %s=%s", column, variable)
	row := a.db.QueryRow(query)

	err := row.Scan(&has)
	if err == nil {
		return fmt.Sprintf("user with this %s: %s already registered", column, variable)
	}
	return ""
}

func (a *AuthDB) CheckPassByEmail(email, pass string) (string, error) {
	row := a.db.QueryRow(`select hash_password, user_id from "auth" where email = $1`, email)

	var corrPass string
	var id int

	err := row.Scan(&corrPass, &id)
	if err != nil {
		return "no user with this email", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(corrPass), []byte(pass)); err != nil {
		return "incorrect password", err
	}

	token, err := createJWTToken(id)
	if err != nil {
		return "server couldn't create JWT token", err
	}

	return token, nil
}

func (a *AuthDB) CheckPassByName(name, pass string) (string, error) {
	row := a.db.QueryRow(`select password, userId from "auth" where name = $1`, name)

	var corrPass string
	var id int

	err := row.Scan(&corrPass, &id)
	if err != nil {
		return "no user with this name", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(corrPass), []byte(pass)); err != nil {
		return "incorrect password", err
	}

	token, err := createJWTToken(id)
	if err != nil {
		return "server couldn't create JWT token", err
	}

	return token, nil
}

func (a *AuthDB) CheckPassByPhone(phone, pass string) (string, error) {
	row := a.db.QueryRow(`select password, userId from "auth" where phone = $1`, phone)

	var corrPass string
	var id int
	err := row.Scan(&corrPass, &id)
	if err != nil {
		return "no user with this phone", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(corrPass), []byte(pass)); err != nil {
		return "incorrect password", err
	}

	token, err := createJWTToken(id)
	if err != nil {
		return "server couldn't create JWT token", err
	}

	return token, nil
}

func createJWTToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func (a *AuthDB) GetUserByID(userId int) (models.User, error) {
	fmt.Println(userId)
	row := a.db.QueryRow(`select * from "user" where user_id = $1`, userId)

	user := models.User{}
	err := row.Scan(&user.UserId, &user.Name, &user.HashPass, &user.Email, &user.Phone, &user.AvatarURL, &user.CreatedAt, &user.UpdatedAt, &user.LastSeen)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
