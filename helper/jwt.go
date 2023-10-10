package helper

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(mapClaims jwt.MapClaims, secretSign []byte) (string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secretSign)
	if err != nil {
		return "",err
	}

	return tokenString,nil
}

func ParseJWT(tokenString string, secretSign []byte) (jwt.MapClaims,error){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")

		return secretSign, nil
	})

	if err != nil {
		return jwt.MapClaims{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims,nil
	}

	return jwt.MapClaims{},errors.New("invalid token")
}