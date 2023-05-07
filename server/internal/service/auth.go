package service

import (
	"github.com/IskanderSh/chat-app/internal"
	"github.com/IskanderSh/chat-app/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

var (
	validName  = regexp.MustCompile(`[a-zA-Z]{2,20}`)
	validLogin = regexp.MustCompile(`[\w*!@#$%^&?]{3,30}`)
	validEmail = regexp.MustCompile(`\w{4,15}@\w{4,8}\.\w{2,5}`)
	validPhone = regexp.MustCompile(`8[0-9]{10}`)
	validPass  = regexp.MustCompile(`[\w*!@#$%^&?]{8,30}`)
)

type IAuthBLogic interface {
	SignUp(ctx *gin.Context, inp internal.UserSignUpInput) string
	SignIn(ctx *gin.Context, inp internal.UserSignInInput) string
}

type AuthBLogic struct {
	database repository.IAuthDB
}

func NewAuthBLogic() *AuthBLogic {
	return &AuthBLogic{
		database: repository.NewAuthDB(),
	}
}

func invalidHandler(ctx *gin.Context, msg string) string {
	internal.NewResponse(ctx, http.StatusBadRequest, "invalid "+msg)
	return ""
}

func (b *AuthBLogic) SignUp(ctx *gin.Context, inp internal.UserSignUpInput) string {
	if !validName.MatchString(inp.Name) {
		return invalidHandler(ctx, "name")
	}
	if !validName.MatchString(inp.Surname) {
		return invalidHandler(ctx, "surname")
	}
	if !validLogin.MatchString(inp.Login) {
		return invalidHandler(ctx, "login")
	}
	if !validEmail.MatchString(inp.Email) {
		return invalidHandler(ctx, "email")
	}
	if !validPhone.MatchString(inp.Phone) {
		return invalidHandler(ctx, "phone")
	}
	if !validPass.MatchString(inp.Password) {
		return invalidHandler(ctx, "password")
	}

	jwtToken, err := b.database.GenerateJWTToken(inp.Email, inp.Password)
	if err != nil {
		internal.NewResponse(ctx, http.StatusInternalServerError, "")
	}

	return jwtToken
}

func (b *AuthBLogic) SignIn(ctx *gin.Context, inp internal.UserSignInInput) string {
	if !validEmail.MatchString(inp.Email) {
		return invalidHandler(ctx, "email")
	}
	if !validPass.MatchString(inp.Password) {
		return invalidHandler(ctx, "password")
	}

	b.database.CheckPass(inp.Email, inp.Password)
}
