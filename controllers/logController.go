package controllers

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tayron/integra-sistema/bootstrap/library/template"
	"github.com/tayron/integra-sistema/models"
)

// ListarLog -
func ListarLog(w http.ResponseWriter, r *http.Request) {

	parametrosURL := mux.Vars(r)
	idIntegracao, _ := strconv.ParseInt(parametrosURL["id"], 10, 64)

	integracao := models.Integracao{}
	log := models.Log{}

	parametros := struct {
		NomeSistema   string
		VersaoSistema string
		Mensagem      string
		Sucesso       bool
		Erro          bool
		Integracao    models.Integracao
		ListaLogs     []models.Log
	}{
		NomeSistema:   os.Getenv("NOME_SISTEMA"),
		VersaoSistema: os.Getenv("VERSAO_SISTEMA"),
		Integracao:    integracao.BuscarPorID(idIntegracao),
		ListaLogs:     log.BuscarPorIDIntegracao(idIntegracao),
	}

	template.LoadView(w, "template/log/*.html", "listarLogPage", parametros)
}
