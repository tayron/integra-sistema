package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/tayron/integra-sistema/models"
)

// ListarIntegracao -
func ListarIntegracao(w http.ResponseWriter, r *http.Request) {

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
	templates.ExecuteTemplate(w, "listarIntegracoesPage", parametros)
}

// CriarIntegracao -
func CriarIntegracao(w http.ResponseWriter, r *http.Request) {
	integracao := models.Integracao{
		Nome:                 r.FormValue("nome"),
		Endpoint:             r.FormValue("endpoint"),
		NomeSistemaOrigem:    r.FormValue("nome_sistema_origem"),
		APISistemaOrigem:     r.FormValue("api_sistema_origem"),
		MetodoSistemaOrigem:  r.FormValue("metodo_sistema_origem"),
		NomeSistemaDestino:   r.FormValue("nome_sistema_destino"),
		APISistemaDestino:    r.FormValue("api_sistema_destino"),
		MetodoSistemaDestino: r.FormValue("metodo_sistema_destino"),
	}

	retornoGravacao := integracao.Gravar()

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
	templates.ExecuteTemplate(w, "homePage", parametros)
}

// EditarIntegracao -
func EditarIntegracao(w http.ResponseWriter, r *http.Request) {

	parametrosURL := mux.Vars(r)
	idIntegracao, _ := strconv.ParseInt(parametrosURL["id"], 10, 64)

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
	templates.ExecuteTemplate(w, "editarIntegracoesPage", parametros)
}

// ExcluirIntegracao -
func ExcluirIntegracao(w http.ResponseWriter, r *http.Request) {
	idIntegracao, _ := strconv.Atoi(r.FormValue("id"))

	integracao := models.Integracao{
		ID: idIntegracao,
	}

	retornoGravacao := integracao.Excluir()

	var mensagem string
	var sucesso bool
	var erro bool

	if retornoGravacao == true {
		sucesso = true
		mensagem = fmt.Sprint("Sucesso ao excluir a integração")
	} else {
		erro = true
		mensagem = fmt.Sprint("Erro ao excluir a integração")
	}

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
		Mensagem:         mensagem,
		Sucesso:          sucesso,
		Erro:             erro,
	}

	var templates = template.Must(template.ParseGlob("template/*.html"))
	template.Must(templates.ParseGlob("template/layout/*.html"))
	template.Must(templates.ParseGlob("template/integracao/*.html"))
	templates.ExecuteTemplate(w, "listarIntegracoesPage", parametros)
}

// SalvarIntegracao -
func SalvarIntegracao(w http.ResponseWriter, r *http.Request) {
	idIntegracao, _ := strconv.Atoi(r.FormValue("id"))

	integracao := models.Integracao{
		ID:                   idIntegracao,
		Nome:                 r.FormValue("nome"),
		Endpoint:             r.FormValue("endpoint"),
		NomeSistemaOrigem:    r.FormValue("nome_sistema_origem"),
		APISistemaOrigem:     r.FormValue("api_sistema_origem"),
		MetodoSistemaOrigem:  r.FormValue("metodo_sistema_origem"),
		NomeSistemaDestino:   r.FormValue("nome_sistema_destino"),
		APISistemaDestino:    r.FormValue("api_sistema_destino"),
		MetodoSistemaDestino: r.FormValue("metodo_sistema_destino"),
	}

	retornoGravacao := integracao.Atualizar()

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
		Integracao    models.Integracao
	}{
		NomeSistema:   os.Getenv("NOME_SISTEMA"),
		VersaoSistema: os.Getenv("VERSAO_SISTEMA"),
		Sucesso:       sucesso,
		Erro:          erro,
		Mensagem:      mensagem,
		Integracao:    integracao,
	}

	var templates = template.Must(template.ParseGlob("template/*.html"))
	template.Must(templates.ParseGlob("template/layout/*.html"))
	template.Must(templates.ParseGlob("template/integracao/*.html"))
	templates.ExecuteTemplate(w, "editarIntegracoesPage", parametros)
}
