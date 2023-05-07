package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/udborets/chat-app/server/internal/responses"
	"github.com/udborets/chat-app/server/internal/service"
	"github.com/udborets/chat-app/server/internal/utilities"
	"net/http"
)

type AuthHTTP struct {
	authBLogic service.IAuthBLogic
}

func NewHTTP(config string) *AuthHTTP {
	return &AuthHTTP{
		authBLogic: service.NewAuthBLogic(config),
	}
}

func (h *AuthHTTP) Start() {
	app := gin.Default()

	authAPI := app.Group("/auth")

	authAPI.POST("/signup", h.userSignUp)
	//authAPI.POST("/singin", h.userSignIn)

	app.Run(":1773")
}

func (h *AuthHTTP) userSignUp(ctx *gin.Context) {
	var inp utilities.UserSignUpInput

	if err := ctx.BindJSON(&inp); err != nil {
		responses.NewResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	h.authBLogic.SignUp(ctx, inp)
}

//func (h *AuthHTTP) userSignIn(ctx *gin.Context) {
//	var inp internal.UserSignInInput
//
//	if err := ctx.BindJSON(&inp); err != nil {
//		internal.NewResponse(ctx, http.StatusBadRequest, "invalid input body")
//		return
//	}
//
//	h.authBLogic.SignIn(ctx, inp)
//}
