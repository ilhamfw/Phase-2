package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func main() {
	// Buat klaim JWT
	claims := jwt.MapClaims{
		"username": "Tugas Meilyanto",
		"role":     "admin",
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	// Buat token JWT dengan klaim di atas
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte("12345")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Gagal membuat token:", err)
		return
	}

	fmt.Println("Token JWT yang dibuat:")
	fmt.Println(tokenString)

	// Parse dan verifikasi token JWT
	parsedToken, parseErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verifikasi bahwa algoritma yang digunakan adalah HMAC-SHA256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Algoritma tanda tangan tidak valid")
		}
		return secretKey, nil
	})

	if parseErr != nil {
		fmt.Println("Gagal mengurai token:", parseErr)
		return
	}

	// Cek apakah token valid
	if parsedToken.Valid {
		fmt.Println("Token JWT valid")
		// Anda dapat mengakses klaim dengan cara berikut:
		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
			fmt.Println("Username:", claims["username"])
			fmt.Println("Role:", claims["role"])
		}
	} else {
		fmt.Println("Token JWT tidak valid")
	}
}
