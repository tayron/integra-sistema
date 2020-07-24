package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tayron/integra-sistema/models"
)

// ProcessarIntegracao -
func ProcessarIntegracao(w http.ResponseWriter, r *http.Request) {
	parametrosURL := mux.Vars(r)
	endpoint := parametrosURL["endpoint"]

	integracao := models.Integracao{}
	integracao = integracao.BuscarPorEndpoint(endpoint)

	sucesso := false
	mensagem := ""

	if integracao.ID == 0 {
		mensagem = "Integração inexistente"

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

	parametroModel := models.Parametro{}
	id := int64(integracao.ID)
	listaParametro := parametroModel.BuscarPorIDIntegracao(id)

	fmt.Println(listaParametro)

	mensagem = fmt.Sprintf("Integração encontrada de ID: %d", integracao.ID)

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
