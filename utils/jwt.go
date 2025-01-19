package utils

import (
	"errors"
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

func VerifyToken(token string) (int64, error) {
	parsedtoken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error){
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil{
		return 0, errors.New("Could not parse token")
	}
	tokenIsValid := parsedtoken.Valid
	if !tokenIsValid {
		return 0, errors.New("could not validate token")
	}

	claims, ok := parsedtoken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("Invalid token claims")
	}

	// email := claims["emails"].(string)
	userId := int64(claims["userID"].(float64))
	return userId, nil
}
