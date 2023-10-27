package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	claims := jwt.MapClaims{
		"username": "Ilham Fw",
		"role":    "admin",
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}

	secretKey := []byte("12345")
	// Method NewWithClaims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}

	fmt.Println("JWT Token:", tokenString)

	// Parse token
	parsedToken, parserErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Algoritma tidak valid")
		}
		return secretKey, nil
	})

	if parserErr != nil {
		fmt.Println("Error while decode token: ", parserErr)
	}

	if parsedToken.Valid {
		fmt.Println("Token is valid")
		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
			fmt.Println("username:", claims["username"])
			fmt.Println("role:", claims["role"])
		}
	}
}
