package template

import (
	"net/http"
	"os"
	"text/template"

	"github.com/tayron/integra-sistema/bootstrap/library/session"
)

var templates = template.Must(template.ParseGlob("template/*.html"))

// ObterInformacaoSistema -
func ObterInformacaoSistema(w http.ResponseWriter, r *http.Request) System {
	nome := session.GetDadoSessao("login", w, r)

	return System{
		Name:    os.Getenv("NOME_SISTEMA"),
		Version: os.Getenv("VERSAO_SISTEMA"),
		Usuario: nome,
	}
}

// LoadView -
func LoadView(w http.ResponseWriter, adicionalPath string, viewName string, parametros interface{}) {
	template.Must(templates.ParseGlob("template/layout/*.html"))

	if adicionalPath != "" {
		template.Must(templates.ParseGlob(adicionalPath))
	}

	templates.ExecuteTemplate(w, viewName, parametros)
}
