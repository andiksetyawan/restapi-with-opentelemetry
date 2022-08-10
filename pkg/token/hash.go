package token

import (
	"golang.org/x/crypto/bcrypt"
)

func HashBuilder(s string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashBytes), nil
}

func HashIsMatch(plain, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	if err != nil {
		return err
	}
	return nil
}
