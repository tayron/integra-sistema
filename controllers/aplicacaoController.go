package controllers

import (
	"net/http"

	appTemplate "github.com/tayron/integra-sistema/bootstrap/library/template"
)

// IndexApplication -
func IndexApplication(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	parametros := appTemplate.Parametro{
		System: appTemplate.ObterInformacaoSistema(w, r),
	}

	appTemplate.LoadView(w, "", "homePage", parametros)
}
