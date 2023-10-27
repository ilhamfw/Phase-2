package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/dgrijalva/jwt-go"
    "github.com/julienschmidt/httprouter"
)

// Middleware JWTAuthentication untuk memeriksa token
func JWTAuthentication(next httprouter.Handle) httprouter.Handle {
    return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        tokenString := r.Header.Get("Authorization")
        if tokenString == "" {
            http.Error(w, "Token missing", http.StatusUnauthorized)
            return
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("Algoritma tidak valid")
            }
            return []byte("12345"), nil
        })

        if err != nil {
            http.Error(w, "Token is invalid", http.StatusUnauthorized)
            return
        }

        if token.Valid {
            next(w, r, ps)
        } else {
            http.Error(w, "Token is invalid", http.StatusUnauthorized)
        }
    }
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    w.Write([]byte("Ini adalah endpoint yang dilindungi"))
}

func main() {
    router := httprouter.New()

    // Endpoint yang dilindungi dengan middleware JWT
    router.GET("/protected", JWTAuthentication(ProtectedEndpoint))

    // Buat token
    claims := jwt.MapClaims{
        "username": "Ilham Fw",
        "role":    "admin",
        "exp":     time.Now().Add(time.Hour * 1).Unix(),
    }

    secretKey := []byte("12345")
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        fmt.Println("Error generating token:", err)
        return
    }

    fmt.Println("JWT Token:", tokenString)

    http.Handle("/", router)
    http.ListenAndServe(":8080", nil)
}
