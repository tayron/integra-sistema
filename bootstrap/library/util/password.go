package util

import (
	"golang.org/x/crypto/bcrypt"
)

// GerarHashSenha -
func GerarHashSenha(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CompararSenhaComHash -
func CompararSenhaComHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

/*
func exemploUso() {
	password := "secret"
	hash, _ := GerarHashSenha(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CompararSenhaComHash(password, hash)
	fmt.Println("Match:   ", match)
}
*/
