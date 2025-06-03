package auth

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

var JwtSecret = []byte("secret-key")

type User struct {
	Login    string
	Password string
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func generateSimplePassword(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	password := make([]byte, length)
	for i := range password {
		password[i] = chars[rand.Intn(len(chars))]
	}
	return string(password)
}

func Generate(id int) (string, string) {
	stringId := fmt.Sprintf("u%04d", id)
	pwd := generateSimplePassword(8)

	return stringId, pwd
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authToken, err := GetAuthToken(r)
		if err != nil {
			next(w, r)
			return
		}

		claims := &Claims{}

		token, err := jwt.ParseWithClaims(authToken, claims, func(token *jwt.Token) (interface{}, error) {
			return JwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Redirect(w, r, "http://u69196.kubsu-dev.ru/login", http.StatusSeeOther)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
		next(w, r.WithContext(ctx))
	}
}
