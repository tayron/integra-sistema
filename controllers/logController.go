package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	appTemplate "github.com/tayron/integra-sistema/bootstrap/library/template"
	"github.com/tayron/integra-sistema/models"
)

// ListarLog -
func ListarLog(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	parametrosURL := mux.Vars(r)
	idIntegracao, _ := strconv.ParseInt(parametrosURL["id"], 10, 64)
	flashMessage := appTemplate.FlashMessage{}

	integracao := models.Integracao{}
	log := models.Log{}

	var Logs = struct {
		Integracao models.Integracao
		ListaLogs  []models.Log
	}{
		Integracao: integracao.BuscarPorID(idIntegracao),
		ListaLogs:  log.BuscarPorIDIntegracao(idIntegracao),
	}

	parametros := appTemplate.Parametro{
		System:       appTemplate.ObterInformacaoSistema(w, r),
		FlashMessage: flashMessage,
		Parametro:    Logs,
	}

	appTemplate.LoadView(w, "template/log/*.html", "listarLogPage", parametros)
}
