package routes

import (
	"net/http"

	"example.com/models"
	"example.com/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context){
	var user models.User
	err := context.ShouldBindJSON(&user)
    if err != nil{
        context.JSON(http.StatusBadRequest, gin.H{"message":"Cant parse Data!"})
        return
    }

	err = user.Save()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Cant save user!"})
        return
	}
}

func login(context *gin.Context){
	var user models.User
	err := context.ShouldBindJSON(&user)
    if err != nil{
        context.JSON(http.StatusBadRequest, gin.H{"message":"Cant parse Data!"})
        return
    }
	
	err = user.ValidateCredentials()
	if err != nil{
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unable to authenticate user"})
		return
	}
	
	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Cant authenticate user!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token" : token})
}
