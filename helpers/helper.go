package helpers

import "golang.org/x/crypto/bcrypt"

const salt = 8

func HashPass(p string) (string, error) {
	password := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(password, salt)
	return string(hash), err
}

func ComparePass(dbPass, userPass string) bool {
	hash, pass := []byte(dbPass), []byte(userPass)

	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}
