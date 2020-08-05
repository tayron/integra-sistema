package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tayron/integra-sistema/bootstrap/library/template"
	"github.com/tayron/integra-sistema/models"
)

// ListarLog -
func ListarLog(w http.ResponseWriter, r *http.Request) {

	parametrosURL := mux.Vars(r)
	idIntegracao, _ := strconv.ParseInt(parametrosURL["id"], 10, 64)
	flashMessage := template.FlashMessage{}

	integracao := models.Integracao{}
	log := models.Log{}

	var Logs = struct {
		Integracao models.Integracao
		ListaLogs  []models.Log
	}{
		Integracao: integracao.BuscarPorID(idIntegracao),
		ListaLogs:  log.BuscarPorIDIntegracao(idIntegracao),
	}

	parametros := template.Parametro{
		System:       template.ObterInformacaoSistema(),
		FlashMessage: flashMessage,
		Parametro:    Logs,
	}

	template.LoadView(w, "template/log/*.html", "listarLogPage", parametros)
}
