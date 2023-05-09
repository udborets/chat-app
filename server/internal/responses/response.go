package responses

import (
	"github.com/gin-gonic/gin"
	"log"
)

type response struct {
	Message string `json:"message"`
}

func NewResponse(ctx *gin.Context, statusCode int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}
	ctx.AbortWithStatusJSON(statusCode, response{msg})
}
