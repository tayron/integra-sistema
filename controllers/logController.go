package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tayron/gopaginacao"
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
	logModel := models.Log{}

	numeroTotalRegistro := logModel.ObterNumeroLogsPorIDIntegracao(idIntegracao)
	htmlPaginacao, offset, err := gopaginacao.CriarPaginacao(numeroTotalRegistro, r)

	var listaLogs []models.Log

	if err == nil {
		listaLogs = logModel.BuscarPorIDIntegracao(idIntegracao, offset)
	}

	var Logs = struct {
		Integracao models.Integracao
		ListaLogs  []models.Log
		Paginacao  template.HTML
	}{
		Integracao: integracao.BuscarPorID(idIntegracao),
		ListaLogs:  listaLogs,
		Paginacao:  template.HTML(htmlPaginacao),
	}

	parametros := appTemplate.Parametro{
		System:       appTemplate.ObterInformacaoSistema(w, r),
		FlashMessage: flashMessage,
		Parametro:    Logs,
	}

	appTemplate.LoadView(w, "template/log/*.html", "listarLogPage", parametros)
}
