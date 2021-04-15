package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JWTSecret = []byte("sTRzrhe9XBR6BbBrX7PJIRS5obO9kd6Gs2T4h2VDHRrUTZqnF0avgVf2AnwzBNtPqZcKsinMk5qJbyZwMhmo0Czs8DvvYRjS2ecJ3ECeZWyhOw4KMSnh7dP6Jjz9D9sg")

func GenerateJWT(email string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	t, _ := token.SignedString(JWTSecret)
	return t
}
