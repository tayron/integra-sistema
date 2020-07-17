package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
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

// ListarHandler -+
func ListarHandler(w http.ResponseWriter, r *http.Request) {

	integracao := models.Integracao{}

	parametros := struct {
		NomeSistema      string
		VersaoSistema    string
		Mensagem         string
		Sucesso          bool
		Erro             bool
		ListaIntegracoes []models.Integracao
	}{
		NomeSistema:      os.Getenv("NOME_SISTEMA"),
		VersaoSistema:    os.Getenv("VERSAO_SISTEMA"),
		ListaIntegracoes: integracao.BuscarTodos(),
	}

	var templates = template.Must(template.ParseGlob("template/*.html"))
	template.Must(templates.ParseGlob("template/layout/*.html"))
	template.Must(templates.ParseGlob("template/integracao/*.html"))
	err := templates.ExecuteTemplate(w, "listarIntegracoesPage", parametros)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// EditarHandler -+
func EditarHandler(w http.ResponseWriter, r *http.Request) {

	parametrosURL := mux.Vars(r)
	idIntegracao, _ := strconv.Atoi(parametrosURL["id"])

	integracao := models.Integracao{}

	parametros := struct {
		NomeSistema   string
		VersaoSistema string
		Mensagem      string
		Sucesso       bool
		Erro          bool
		Integracao    models.Integracao
	}{
		NomeSistema:   os.Getenv("NOME_SISTEMA"),
		VersaoSistema: os.Getenv("VERSAO_SISTEMA"),
		Integracao:    integracao.BuscarPorID(idIntegracao),
	}

	var templates = template.Must(template.ParseGlob("template/*.html"))
	template.Must(templates.ParseGlob("template/layout/*.html"))
	template.Must(templates.ParseGlob("template/integracao/*.html"))
	err := templates.ExecuteTemplate(w, "editarIntegracoesPage", parametros)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
