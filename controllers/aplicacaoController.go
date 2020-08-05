package controllers

import (
	"net/http"

	"github.com/tayron/integra-sistema/bootstrap/library/template"
)

// IndexApplication -
func IndexApplication(w http.ResponseWriter, r *http.Request) {
	parametros := template.Parametro{
		System: template.ObterInformacaoSistema(),
	}

	template.LoadView(w, "", "homePage", parametros)
}
