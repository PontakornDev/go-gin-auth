package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(secretKey, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(secretKey))
	return err == nil
}
