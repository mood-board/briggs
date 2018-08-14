package utils

import (
	"golang.org/x/crypto/bcrypt"
)

//HashPassword Uses Bcrypt to hash a users password
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//CheckPassword Compares password and hash for equality
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
