package handlers

import (
	"net/http"
	"os"
	"text/template"
)

// InicioHandler Controlador página inicial
func InicioHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

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

	err = tmpl.Execute(w, parametros)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
