package controllers

import (
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
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

	var templates = template.Must(template.ParseGlob("template/*.html"))
	template.Must(templates.ParseGlob("template/layout/*.html"))
	template.Must(templates.ParseGlob("template/log/*.html"))
	templates.ExecuteTemplate(w, "listarLogPage", parametros)
}
