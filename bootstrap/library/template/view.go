package template

import (
	"net/http"
	"os"
	"text/template"
)

var templates = template.Must(template.ParseGlob("template/*.html"))

// LoadView -
func LoadView(w http.ResponseWriter, adicionalPath string, viewName string, parametros interface{}) {

	template.Must(templates.ParseGlob("template/layout/*.html"))

	if adicionalPath != "" {
		template.Must(templates.ParseGlob(adicionalPath))
	}

	templates.ExecuteTemplate(w, viewName, parametros)
}

// ObterSystemInformation -
func ObterSystemInformation() System {
	return System{
		Name:    os.Getenv("NOME_SISTEMA"),
		Version: os.Getenv("VERSAO_SISTEMA"),
	}
}

type System struct {
	Name    string
	Version string
}

type FlashMessage struct {
	Type    string
	Message string
}

type Parameter struct {
	System       System
	FlashMessage FlashMessage
	Parameter    interface{}
}
