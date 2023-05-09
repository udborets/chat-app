package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/udborets/chat-app/server/internal/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"regexp"
	"time"
)

type IAuthBLogic interface {
	SignUp(inp models.UserSignUpInput) (int, string, error)
	SignIn(inp models.UserSignInInput) (int, string, error)
	ParseJWTToken(stringToken string) (interface{}, int, string, error)
}

type AuthBLogic struct {
	database IAuthDB
}

func NewAuthBLogic(config string) *AuthBLogic {
	return &AuthBLogic{
		database: NewAuthDB(config),
	}
}

var (
	validName  = regexp.MustCompile(`[a-zA-Zа-яА-ЯёЁ]{2,20}`)
	validEmail = regexp.MustCompile(`\w{4,15}@\w{4,8}\.\w{2,5}`)
	validPhone = regexp.MustCompile(`[0-9]{8,15}`)
	validPass  = regexp.MustCompile(`[\w*!@#$%^&?]{8,30}`)
)

func (b *AuthBLogic) SignUp(inp models.UserSignUpInput) (int, string, error) {
	if !validName.MatchString(inp.Name) && !validPhone.MatchString(inp.Phone) && !validEmail.MatchString(inp.Email) {
		return http.StatusBadRequest, "invalid login", errors.New("invalid login")
	} else if !validName.MatchString(inp.Name) {
		return http.StatusBadRequest, "invalid name", errors.New("invalid name")
	} else if !validEmail.MatchString(inp.Email) {
		return http.StatusBadRequest, "invalid email", errors.New("invalid email")
	} else if inp.Phone != "" && !validPhone.MatchString(inp.Phone) {
		return http.StatusBadRequest, "invalid phone", errors.New("invalid phone")
	}
	if !validPass.MatchString(inp.Password) {
		return http.StatusBadRequest, "invalid password", errors.New("invalid password")
	}

	msg, err := b.database.CheckUniqUser(inp)
	if err != nil {
		return http.StatusBadRequest, msg, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(inp.Password), 10)
	if err != nil {
		return http.StatusInternalServerError, "server couldn't hash password", err
	}

	user := models.User{
		Name:      inp.Name,
		Email:     inp.Email,
		Phone:     inp.Phone,
		HashPass:  string(hash),
		AvatarURL: inp.AvatarURL,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		LastSeen:  time.Now().Unix(),
	}

	msg, err = b.database.AddUser(user)
	if err != nil {
		return http.StatusInternalServerError, msg, err
	}
	return http.StatusOK, "successful signup", errors.New("successful signup")
}

func (b *AuthBLogic) SignIn(inp models.UserSignInInput) (int, string, error) {
	var jwtToken string
	if validEmail.MatchString(inp.Login) {
		msg, err := b.database.CheckPassByEmail(inp.Login, inp.Password)
		if err != nil {
			return http.StatusUnauthorized, msg, err
		}
		jwtToken = msg
	} else if validName.MatchString(inp.Login) {
		msg, err := b.database.CheckPassByName(inp.Login, inp.Password)
		if err != nil {
			return http.StatusUnauthorized, msg, err
		}
		jwtToken = msg
	} else if validPhone.MatchString(inp.Login) {
		msg, err := b.database.CheckPassByPhone(inp.Login, inp.Password)
		if err != nil {
			return http.StatusUnauthorized, msg, err
		}
		jwtToken = msg
	}

	return http.StatusOK, jwtToken, nil
}

func (b *AuthBLogic) ParseJWTToken(tokenString string) (interface{}, int, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return nil, http.StatusUnauthorized, "server couldn't parse JWT token", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return nil, http.StatusUnauthorized, "expiration of JWT ended", nil
		}

		user, err := b.database.GetUserByID(int(claims["sub"].(float64)))
		if err != nil {
			return nil, http.StatusBadRequest, "no user with this id", err
		}

		return user, http.StatusOK, "", nil
	} else {
		return nil, http.StatusUnauthorized, "JWT token has error", errors.New("JWT token has error")
	}
}
