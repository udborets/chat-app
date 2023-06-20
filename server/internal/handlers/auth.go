package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/udborets/chat-app/server/internal/models"
	"github.com/udborets/chat-app/server/internal/responses"
	"github.com/udborets/chat-app/server/internal/service"
	"net/http"
)

type AuthHandler struct {
	authService service.IAuthService
}

func NewHTTP() *AuthHandler {
	return &AuthHandler{
		authService: service.NewAuthService(),
	}
}

func (h *AuthHandler) InitRoutes(router *gin.Engine) {
	authAPI := router.Group("/auth")

	authAPI.POST("/signup", h.userSignUp)
	authAPI.POST("/signin", h.userSignIn)
	authAPI.GET("/validate", h.requireAuth, h.validate)
}

func (h *AuthHandler) userSignUp(ctx *gin.Context) {
	var inp models.UserSignUpInput

	if err := ctx.BindJSON(&inp); err != nil {
		responses.NewResponse(ctx, http.StatusBadRequest, "invalid input body", err)
		return
	}

	statusCode, msg, err := h.authService.SignUp(inp)
	responses.NewResponse(ctx, statusCode, msg, err)
}

func (h *AuthHandler) userSignIn(ctx *gin.Context) {
	var inp models.UserSignInInput

	if err := ctx.BindJSON(&inp); err != nil {
		responses.NewResponse(ctx, http.StatusBadRequest, "invalid input body", err)
		return
	}

	statusCode, msg, err := h.authService.SignIn(inp)
	if err != nil {
		responses.NewResponse(ctx, statusCode, msg, err)
		return
	}

	jwtToken := msg
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", jwtToken, 3600*24*30, "", "", false, true)
	ctx.Redirect(http.StatusMovedPermanently, "http://localhost/auth/validate")
}

func (h *AuthHandler) requireAuth(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("Authorization")
	if err != nil {
		responses.NewResponse(ctx, http.StatusUnauthorized, "Please login to chat or sign in if you first time here)", err)
		return
	}

	output, statusCode, msg, err := h.authService.ParseJWTToken(tokenString)
	if err != nil {
		responses.NewResponse(ctx, statusCode, msg, err)
		return
	}

	ctx.Set("output", output.(models.ValidateOutput))
	ctx.Next()
}

func (h *AuthHandler) validate(ctx *gin.Context) {
	output, _ := ctx.Get("output")
	ctx.JSON(http.StatusOK, output)
}
