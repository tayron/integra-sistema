package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// ProcessarIntegracao -
func ProcessarIntegracao(w http.ResponseWriter, r *http.Request) {
	parametrosURL := mux.Vars(r)
	endpoint := parametrosURL["endpoint"]

	sucesso := true
	mensagem := "Integracao feira com sucesso: " + endpoint

	parametros := struct {
		Sucesso  bool
		Mensagem string
	}{
		Sucesso:  sucesso,
		Mensagem: mensagem,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(parametros)
}
