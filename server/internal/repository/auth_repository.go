package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
	"github.com/udborets/chat-app/server/internal/models"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type IAuthDB interface {
	AddUser(user models.User) (string, error)
	CheckUniqUser(user models.UserSignUpInput) (string, error)
	CheckPassByName(name, pass string) (string, error)
	CheckPassByEmail(email, pass string) (string, error)
	CheckPassByPhone(email, pass string) (string, error)
	GetInfoByUserId(userId int) (interface{}, error)
	//GetUserByID(userId int) (interface{}, error)
	//GetChatsByUserID(userId int) (interface{}, error)
	//GetMessagesByChatsID(chatsId []int) (interface{}, error)
}

type AuthDB struct {
	db *sql.DB
}

func NewAuthDB() *AuthDB {
	return &AuthDB{db: database}
}

func (a *AuthDB) AddUser(user models.User) (string, error) {
	var addedId int

	err := a.db.QueryRow("insert into \"users\" (name, hash_password, email, phone, avatar_url, created_at, updated_at) "+
		"values ($1,$2,$3,$4,$5,$6,$7) returning id", user.Name, user.HashPass, user.Email, user.Phone, user.AvatarURL,
		user.CreatedAt, user.UpdatedAt).Scan(&addedId)
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
	var user_id int

	if user.Name != "" {
		row := a.db.QueryRow("select user_id from \"auth\" where name=$1", user.Name)
		row.Scan(&user_id)
		if user_id != 0 {
			return fmt.Sprintf("user with this name: %s already registered", user.Name), errors.New("server couldn't add row to unique column")
		}
	}

	if user.Email != nil {
		row := a.db.QueryRow("select user_id from \"auth\" where email=$1", user.Email)
		row.Scan(&user_id)
		if user_id != 0 {
			return fmt.Sprintf("user with this email: %s already registered", user.Email), errors.New("server couldn't add row to unique column")
		}
	}

	if user.Phone != nil {
		row := a.db.QueryRow("select user_id from \"auth\" where phone=$1", user.Phone)
		row.Scan(&user_id)
		if user_id != 0 {
			return fmt.Sprintf("user with this phone: %s already registered", user.Phone), errors.New("server couldn't add row to unique column")
		}
	}
	return "", nil
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
	row := a.db.QueryRow(`select hash_password, user_id from "auth" where name = $1`, name)

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
	row := a.db.QueryRow(`select hash_password, user_id from "auth" where phone = $1`, phone)

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

func (a *AuthDB) GetInfoByUserId(userId int) (interface{}, error) {
	output := models.ValidateOutput{}
	selectUser := a.db.QueryRow("select id, name, email, phone, avatar_url from \"users\" where id = $1", userId)

	err := selectUser.Scan(&output.UserId, &output.Name, &output.Email, &output.Phone, &output.AvatarURL)
	if err != nil {
		return nil, err
	}

	chats := make([]models.ChatValidateOutput, 0)
	selectChats, err := a.db.Query("select chat_id from \"users_chats\" where user_id = $1", userId)
	if err != nil {
		return nil, err
	}

	for selectChats.Next() {
		var chat models.ChatValidateOutput
		err := selectChats.Scan(&chat.ChatId)
		if err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}

	for _, chat := range chats {
		chatId := chat.ChatId
		selectMessages, err := a.db.Query("select message_id, text, is_seen from \"messages\" where chat_id = $1", chatId)
		if err != nil {
			return nil, err
		}

		for selectMessages.Next() {
			message := models.MessageValidateOutput{}
			err := selectMessages.Scan(&message.MessageId, &message.Text, &message.IsSeen)
			if err != nil {
				return nil, err
			}

			chat.Messages = append(chat.Messages, message)
		}
		output.Chats = append(output.Chats, chat)
	}
	return output, nil
}
