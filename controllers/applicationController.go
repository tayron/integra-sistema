package controllers

import (
	"net/http"
	"os"
	"text/template"
)

// IndexApplication -
func IndexApplication(w http.ResponseWriter, r *http.Request) {

	parametros := struct {
		NomeSistema   string
		VersaoSistema string
		Mensagem      string
		Sucesso       bool
		Erro          bool
	}{
		NomeSistema:   os.Getenv("NOME_SISTEMA"),
		VersaoSistema: os.Getenv("VERSAO_SISTEMA"),
	}

	var templates = template.Must(template.ParseGlob("template/*.html"))
	template.Must(templates.ParseGlob("template/layout/*.html"))
	templates.ExecuteTemplate(w, "homePage", parametros)
}
