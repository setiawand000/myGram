package helpers

import "golang.org/x/crypto/bcrypt"

func HashPass(pass string) string {
	passH := 8

	hash, _ := bcrypt.GenerateFromPassword([]byte(pass), passH)
	return string(hash)
}

func ComparePass(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}
