package handlers

import (
	"net/http"
	"text/template"
)

// CriarIntegracaoHandler - Grava uma nova integração
func CriarIntegracaoHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data := struct {
		AppVersion string
	}{
		AppVersion: "1.0",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
