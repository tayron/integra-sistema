package handlers

import (
	"net/http"
	"text/template"
)

// InicioHandler Controlador página inicial
func InicioHandler(w http.ResponseWriter, r *http.Request) {
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
