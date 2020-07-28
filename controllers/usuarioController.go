package controllers

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func gerarHashSenha(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func compararSenhaComHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func exemploUso() {
	password := "secret"
	hash, _ := gerarHashSenha(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := compararSenhaComHash(password, hash)
	fmt.Println("Match:   ", match)
}

// ListarUsuario -
func ListarUsuario(w http.ResponseWriter, r *http.Request) {

}

// CadastrarUsuario -
func CadastrarUsuario(w http.ResponseWriter, r *http.Request) {

}

// EditarUsuario -
func EditarUsuario(w http.ResponseWriter, r *http.Request) {

}

// ExcluirUsuario -
func ExcluirUsuario(w http.ResponseWriter, r *http.Request) {

}
