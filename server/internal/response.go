package internal

import (
	"github.com/gin-gonic/gin"
	"log"
)

type response struct {
	Message string `json:"message"`
}

func NewResponse(ctx *gin.Context, statusCode int, msg string) {
	log.Println(msg)
	ctx.AbortWithStatusJSON(statusCode, response{msg})
}
