package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		fmt.Println("jwt validata")
		ctx.Next()
	}
}