package controllers

import (
	"net/http"

	"github.com/tayron/integra-sistema/bootstrap/library/template"
)

// IndexApplication -
func IndexApplication(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	parametros := template.Parametro{
		System: template.ObterInformacaoSistema(w, r),
	}

	template.LoadView(w, "", "homePage", parametros)
}
