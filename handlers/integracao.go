package handlers

import (
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/tayron/integra-sistema/models"
)

// CriarIntegracaoHandler - Grava uma nova integração
func CriarIntegracaoHandler(w http.ResponseWriter, r *http.Request) {
	integracao := models.Integracao{
		Nome:                 r.FormValue("nome"),
		NomeSistemaOrigem:    r.FormValue("nome_sistema_origem"),
		APISistemaOrigem:     r.FormValue("api_sistema_origem"),
		MetodoSistemaOrigem:  r.FormValue("metodo_sistema_origem"),
		NomeSistemaDestino:   r.FormValue("nome_sistema_destino"),
		APISistemaDestino:    r.FormValue("api_sistema_destino"),
		MetodoSistemaDestino: r.FormValue("metodo_sistema_destino"),
	}

	retornoGravacao := integracao.Gravar(integracao)

	var mensagem string
	var sucesso bool
	var erro bool

	if retornoGravacao == true {
		sucesso = true
		mensagem = fmt.Sprint("Sucesso ao gravar dados da integração")
	} else {
		erro = true
		mensagem = fmt.Sprint("Erro ao gravar dados da integração")
	}

	parametros := struct {
		NomeSistema   string
		VersaoSistema string
		Sucesso       bool
		Erro          bool
		Mensagem      string
	}{
		NomeSistema:   os.Getenv("NOME_SISTEMA"),
		VersaoSistema: os.Getenv("VERSAO_SISTEMA"),
		Sucesso:       sucesso,
		Erro:          erro,
		Mensagem:      mensagem,
	}

	var templates = template.Must(template.ParseGlob("template/*.html"))
	template.Must(templates.ParseGlob("template/layout/*.html"))
	err := templates.ExecuteTemplate(w, "homePage", parametros)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}