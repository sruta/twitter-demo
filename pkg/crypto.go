package pkg

import (
	"golang.org/x/crypto/bcrypt"
)

func HashString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	return string(b), err
}

func CheckHashedString(hash string, original string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(original)) == nil
}
