package delivery

import (
	"fmt"
	"github.com/IskanderSh/chat-app/internal"
	"github.com/IskanderSh/chat-app/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHTTP struct {
	authBLogic service.IAuthBLogic
}

func NewHTTP() *AuthHTTP {
	return &AuthHTTP{
		authBLogic: service.NewAuthBLogic(),
	}
}

func (h *AuthHTTP) Start() {
	app := gin.Default()

	authAPI := app.Group("/auth")

	authAPI.POST("/sign-up", h.userSignUp)
	//authAPI.POST("/sing-in", h.userSignIn)

	app.Run(":3000")
}

func (h *AuthHTTP) userSignUp(ctx *gin.Context) {
	var inp internal.UserSignUpInput

	if err := ctx.BindJSON(&inp); err != nil {
		internal.NewResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	jwtToken := h.authBLogic.SignUp(ctx, inp)
	fmt.Println(jwtToken)
}

func (h *AuthHTTP) userSignIn(ctx *gin.Context) {
	var inp internal.UserSignInInput

	if err := ctx.BindJSON(&inp); err != nil {
		internal.NewResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	h.authBLogic.SignIn(ctx, inp)
}
