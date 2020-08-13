package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tayron/integra-sistema/api/services"
	"github.com/tayron/integra-sistema/models"
)

// ProcessarIntegracao -
func ProcessarIntegracao(w http.ResponseWriter, r *http.Request) {
	parametrosURL := mux.Vars(r)
	endpoint := parametrosURL["endpoint"]

	integracao := models.Integracao{}
	integracao = integracao.BuscarPorEndpoint(endpoint)

	if integracao.ID == 0 {
		retornarMensagem("Integracao não encontrada", false, w, r)
	} else {
		processarIntegracao(integracao, w, r)
	}
}

func processarIntegracao(integracao models.Integracao, w http.ResponseWriter, r *http.Request) {
	parametroModel := models.Parametro{}
	id := int64(integracao.ID)
	listaParametros := parametroModel.BuscarPorIDIntegracao(id)

	if len(listaParametros) == 0 {
		var mensagem []byte = []byte("{'status': false, 'mensagem': 'Nenhum parametro configurado para integração'}")
		retornarMensagemResposta(mensagem, false, w, r)
		return
	}

	switch integracao.MetodoSistemaDestino {
	case "POST":
		mensagem, sucesso := services.EnviarRequisicaoViaPOST(integracao, "POST", listaParametros, r)
		retornarMensagemResposta(mensagem, sucesso, w, r)
	case "POST-JSON":
		mensagem, sucesso := services.EnviarRequisicaoJSONViaPOST(integracao, "POST-JSON", listaParametros, r)
		retornarMensagemResposta(mensagem, sucesso, w, r)
	case "GET":
		mensagem, sucesso := services.EnviarRequisicaoViaGET(integracao, "GET", listaParametros, r)
		retornarMensagemResposta(mensagem, sucesso, w, r)
	}
}

func retornarMensagem(mensagem string, status bool, w http.ResponseWriter, r *http.Request) {
	parametros := struct {
		Sucesso  bool
		Mensagem string
	}{
		Sucesso:  status,
		Mensagem: mensagem,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(parametros)
}

func retornarMensagemResposta(mensagem []byte, status bool, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(mensagem)
}
