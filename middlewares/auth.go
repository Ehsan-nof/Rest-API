package middlewares

import (
	"fmt"
	"net/http"

	"example.com/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == ""{
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"NOT authorized"})
		return
	}
	
	fmt.Print("salam")
	userId, err := utils.VerifyToken(token)

	if err != nil{
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"NOT authorized"})
		return
	}

	context.Set("userId", userId)

	context.Next()
}
