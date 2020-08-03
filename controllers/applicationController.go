package controllers

import (
	"net/http"

	"github.com/tayron/integra-sistema/bootstrap/library/template"
)

// IndexApplication -
func IndexApplication(w http.ResponseWriter, r *http.Request) {
	parametros := template.Parameter{
		System: template.ObterSystemInformation(),
	}

	template.LoadView(w, "", "homePage", parametros)
}
