package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(password string) string {
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
