package util

import (
	"golang.org/x/crypto/bcrypt"
)

const PER_PAGE = 30

func GeneratePasswordHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 1)
	return string(bytes)
}

func CheckPassword(password string, password_hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(password))
	return err == nil
}
