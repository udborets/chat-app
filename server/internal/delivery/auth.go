package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/udborets/chat-app/server/internal/models"
	"github.com/udborets/chat-app/server/internal/responses"
	"github.com/udborets/chat-app/server/internal/service"
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
	authAPI.POST("/signin", h.userSignIn)

	app.Run(":1773")
}

func (h *AuthHTTP) userSignUp(ctx *gin.Context) {
	var inp models.UserSignUpInput

	if err := ctx.BindJSON(&inp); err != nil {
		responses.NewResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	h.authBLogic.SignUp(ctx, inp)
}

func (h *AuthHTTP) userSignIn(ctx *gin.Context) {
	var inp models.UserSignInInput

	if err := ctx.BindJSON(&inp); err != nil {
		responses.NewResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	jwtToken := h.authBLogic.SignIn(ctx, inp)
	if jwtToken == "" {
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", jwtToken, 3600*24*30, "", "", false, true)
}
