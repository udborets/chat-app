package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/udborets/chat-app/server/internal/models"
	"github.com/udborets/chat-app/server/internal/repository"
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
	database repository.IAuthDB
}

func NewAuthBLogic() *AuthBLogic {
	return &AuthBLogic{
		database: repository.NewAuthDB(),
	}
}

var (
	validName  = regexp.MustCompile(`[a-zA-Zа-яА-ЯёЁ]{2,20}`)
	validEmail = regexp.MustCompile(`\w{4,15}@\w{4,8}\.\w{2,5}`)
	validPhone = regexp.MustCompile(`[0-9]{8,15}`)
	validPass  = regexp.MustCompile(`[\w*!@#$%^&?]{8,30}`)
)

func (b *AuthBLogic) SignUp(inp models.UserSignUpInput) (int, string, error) {
	if inp.Email == nil && inp.Phone == nil {
		return http.StatusBadRequest, "no email and phone, at least one is required", errors.New("no email and phone, at least one is required")
	}
	if !validName.MatchString(inp.Name) {
		return http.StatusBadRequest, "invalid name", errors.New("invalid name")
	}
	if inp.Email != nil && !validEmail.MatchString(inp.Email.(string)) {
		return http.StatusBadRequest, "invalid email", errors.New("invalid email")
	}
	if inp.Phone != nil && !validPhone.MatchString(inp.Phone.(string)) {
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
	}

	msg, err = b.database.AddUser(user)
	if err != nil {
		return http.StatusInternalServerError, msg, err
	}
	return http.StatusOK, "successful signup", errors.New("successful signup")
}

func (b *AuthBLogic) SignIn(inp models.UserSignInInput) (int, string, error) {
	var jwtToken string
	if inp.Email != "" && validEmail.MatchString(inp.Email) {
		msg, err := b.database.CheckPassByEmail(inp.Email, inp.Password)
		if err != nil {
			return http.StatusUnauthorized, msg, err
		}
		jwtToken = msg
	} else if inp.Name != "" && validName.MatchString(inp.Name) {
		msg, err := b.database.CheckPassByName(inp.Name, inp.Password)
		if err != nil {
			return http.StatusUnauthorized, msg, err
		}
		jwtToken = msg
	} else if inp.Phone != "" && validPhone.MatchString(inp.Phone) {
		msg, err := b.database.CheckPassByPhone(inp.Phone, inp.Password)
		if err != nil {
			return http.StatusUnauthorized, msg, err
		}
		jwtToken = msg
	} else {
		return http.StatusBadRequest, "incorrect input data, check email/name/phone", errors.New("incorrect input data, check email/name/phone")
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
