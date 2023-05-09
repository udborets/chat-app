package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/udborets/chat-app/server/internal/models"
	"github.com/udborets/chat-app/server/internal/repository"
	"github.com/udborets/chat-app/server/internal/responses"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
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
	ParseJWTToken(ctx *gin.Context, stringToken string)
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
		UpdatedAt: time.Now().Unix(),
		LastSeen:  time.Now().Unix(),
	}

	err, msg := b.database.AddUser(user)
	if err != nil {
		responses.NewResponse(ctx, http.StatusInternalServerError, msg)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "add user to database"})
}

func (b *AuthBLogic) SignIn(ctx *gin.Context, inp models.UserSignInInput) string {
	var jwtToken string
	if validEmail.MatchString(inp.Login) {
		err, msg := b.database.CheckPassByEmail(inp.Login, inp.Password)
		if err != nil {
			responses.NewResponse(ctx, http.StatusUnauthorized, msg)
			return ""
		}
		jwtToken = msg
	} else if validName.MatchString(inp.Login) {
		err, msg := b.database.CheckPassByName(inp.Login, inp.Password)
		if err != nil {
			responses.NewResponse(ctx, http.StatusUnauthorized, msg)
			return ""
		}
		jwtToken = msg
	} else if validPhone.MatchString(inp.Login) {
		err, msg := b.database.CheckPassByPhone(inp.Login, inp.Password)
		if err != nil {
			responses.NewResponse(ctx, http.StatusUnauthorized, msg)
			return ""
		}
		jwtToken = msg
	}

	return jwtToken
}

func (b *AuthBLogic) ParseJWTToken(ctx *gin.Context, tokenString string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		responses.NewResponse(ctx, http.StatusUnauthorized, err.Error())
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			responses.NewResponse(ctx, http.StatusUnauthorized, "")
		}

		err, user := b.database.GetUserByID(int(claims["sub"].(float64)))
		if err != nil {
			fmt.Println(err.Error())
			responses.NewResponse(ctx, http.StatusBadRequest, "no user with this id")
		}

		ctx.Set("user", user)

		ctx.Next()
	} else {
		responses.NewResponse(ctx, http.StatusUnauthorized, "")
	}
}
