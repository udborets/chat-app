package service

import (
	"github.com/gin-gonic/gin"
	"github.com/udborets/chat-app/server/internal/models"
	"github.com/udborets/chat-app/server/internal/repository"
	"github.com/udborets/chat-app/server/internal/responses"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"regexp"
	"time"
)

var (
	validName  = regexp.MustCompile(`[a-zA-Zа-яА-ЯёЁ]{2,20}`)
	validEmail = regexp.MustCompile(`\w{4,15}@\w{4,8}\.\w{2,5}`)
	validPhone = regexp.MustCompile(`[0-9]{8,15}`)
	validPass  = regexp.MustCompile(`[\w*!@#$%^&?]{8,30}`)
)

type IAuthBLogic interface {
	SignUp(ctx *gin.Context, inp models.UserSignUpInput)
	SignIn(ctx *gin.Context, inp models.UserSignInInput) string
}

type AuthBLogic struct {
	database repository.IAuthDB
}

func NewAuthBLogic(config string) *AuthBLogic {
	return &AuthBLogic{
		database: repository.NewAuthDB(config),
	}
}

func invalidHandler(ctx *gin.Context, msg string) string {
	responses.NewResponse(ctx, http.StatusBadRequest, "invalid "+msg)
	return ""
}

func (b *AuthBLogic) SignUp(ctx *gin.Context, inp models.UserSignUpInput) {
	if !validName.MatchString(inp.Name) {
		invalidHandler(ctx, "name")
	}
	if !validEmail.MatchString(inp.Email) {
		invalidHandler(ctx, "email")
	}
	if inp.Phone != "" && !validPhone.MatchString(inp.Phone) {
		invalidHandler(ctx, "phone")
	}
	if !validPass.MatchString(inp.Password) {
		invalidHandler(ctx, "password")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(inp.Password), 10)
	if err != nil {
		log.Printf("error hashing password: %s", inp.Password)
		return
	}

	user := models.User{
		Name:      inp.Name,
		Email:     inp.Email,
		Phone:     inp.Phone,
		HashPass:  string(hash),
		AvatarURL: inp.AvatarURL,
		CreatedAt: time.Now().Unix(),
	}

	err, msg := b.database.AddUser(user)
	if err != nil {
		responses.NewResponse(ctx, http.StatusInternalServerError, msg)
		return
	}
	responses.NewResponse(ctx, http.StatusOK, "add user to database")
}

func (b *AuthBLogic) SignIn(ctx *gin.Context, inp models.UserSignInInput) string {
	var jwtToken string
	if validEmail.MatchString(inp.Login) {
		err, msg := b.database.CheckPassByEmail(inp.Login, inp.Password)
		if err != nil {
			responses.NewResponse(ctx, http.StatusBadRequest, msg)
			return ""
		}
		jwtToken = msg
	} else if validName.MatchString(inp.Login) {
		err, msg := b.database.CheckPassByName(inp.Login, inp.Password)
		if err != nil {
			responses.NewResponse(ctx, http.StatusBadRequest, msg)
			return ""
		}
		jwtToken = msg
	} else if validPhone.MatchString(inp.Login) {
		err, msg := b.database.CheckPassByPhone(inp.Login, inp.Password)
		if err != nil {
			responses.NewResponse(ctx, http.StatusBadRequest, msg)
			return ""
		}
		jwtToken = msg
	}
	responses.NewResponse(ctx, http.StatusOK, "jwt token created")

	return jwtToken
}
