package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "seufvysouabvoa12312331eubveuwo"

func GenerateToken(email string, userId int64) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"emaill": email,
		"userID": userId,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})	

	return token.SignedString([]byte(secretKey))
}
