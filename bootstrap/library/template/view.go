package template

import (
	"net/http"
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
