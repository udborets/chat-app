package service

import (
	"github.com/gin-gonic/gin"
	"github.com/udborets/chat-app/server/internal/repository"
	"github.com/udborets/chat-app/server/internal/responses"
	"github.com/udborets/chat-app/server/internal/utilities"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"regexp"
	"time"
)

var (
	validName  = regexp.MustCompile(`[a-zA-Z]{2,20}`)
	validEmail = regexp.MustCompile(`\w{4,15}@\w{4,8}\.\w{2,5}`)
	validPhone = regexp.MustCompile(`8[0-9]{10}`)
	validPass  = regexp.MustCompile(`[\w*!@#$%^&?]{8,30}`)
)

type IAuthBLogic interface {
	SignUp(ctx *gin.Context, inp utilities.UserSignUpInput)
	//SignIn(ctx *gin.Context, inp internal.UserSignInInput) string
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

func (b *AuthBLogic) SignUp(ctx *gin.Context, inp utilities.UserSignUpInput) {
	if !validName.MatchString(inp.Name) {
		invalidHandler(ctx, "name")
	}
	if !validEmail.MatchString(inp.Email) {
		invalidHandler(ctx, "email")
	}
	if !validPhone.MatchString(inp.Phone) {
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

	user := utilities.User{
		Name:      inp.Name,
		Email:     inp.Email,
		Phone:     inp.Phone,
		HashPass:  hash,
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

//func (b *AuthBLogic) SignIn(ctx *gin.Context, inp internal.UserSignInInput) string {
//	if validEmail.MatchString(inp.Login) {
//
//	}
//	if !validPass.MatchString(inp.Password) {
//		return invalidHandler(ctx, "password")
//	}
//
//	b.database.CheckPass(inp.Login, inp.Password)
//}
