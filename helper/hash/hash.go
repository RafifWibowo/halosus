package hash

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	salt, _ := strconv.Atoi(os.Getenv("BCRYPT_SALT"))
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(bytes), err
}

func CheckPassword(password string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}