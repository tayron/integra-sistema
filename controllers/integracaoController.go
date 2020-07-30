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

// CadastrarIntegracao -
func CadastrarIntegracao(w http.ResponseWriter, r *http.Request) {

	var mensagem string
	var sucesso bool
	var erro bool

	if r.Method == "POST" {
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

		if retornoGravacao == true {
			sucesso = true
			mensagem = fmt.Sprint("Sucesso ao gravar dados da integração")
		} else {
			erro = true
			mensagem = fmt.Sprint("Erro ao gravar dados da integração")
		}
	}

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
		Mensagem:         mensagem,
		Sucesso:          sucesso,
		Erro:             erro,
		ListaIntegracoes: integracao.BuscarTodos(),
	}

	var templates = template.Must(template.ParseGlob("template/*.html"))
	template.Must(templates.ParseGlob("template/layout/*.html"))
	template.Must(templates.ParseGlob("template/integracao/*.html"))
	templates.ExecuteTemplate(w, "cadastrarIntegracoesPage", parametros)
}

// EditarIntegracao -
func EditarIntegracao(w http.ResponseWriter, r *http.Request) {

	parametrosURL := mux.Vars(r)
	idIntegracao, _ := strconv.ParseInt(parametrosURL["id"], 10, 64)

	var mensagem string
	var sucesso bool
	var erro bool

	if r.Method == "POST" {
		id, _ := strconv.Atoi(r.FormValue("id"))
		integracaoModel := models.Integracao{
			ID:                   id,
			Nome:                 r.FormValue("nome"),
			Endpoint:             r.FormValue("endpoint"),
			NomeSistemaOrigem:    r.FormValue("nome_sistema_origem"),
			APISistemaOrigem:     r.FormValue("api_sistema_origem"),
			MetodoSistemaOrigem:  r.FormValue("metodo_sistema_origem"),
			NomeSistemaDestino:   r.FormValue("nome_sistema_destino"),
			APISistemaDestino:    r.FormValue("api_sistema_destino"),
			MetodoSistemaDestino: r.FormValue("metodo_sistema_destino"),
		}

		retornoGravacao := integracaoModel.Atualizar()

		if retornoGravacao == true {
			sucesso = true
			mensagem = fmt.Sprint("Sucesso ao gravar dados da integração")
		} else {
			erro = true
			mensagem = fmt.Sprint("Erro ao gravar dados da integração")
		}
	}

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
		Mensagem:      mensagem,
		Sucesso:       sucesso,
		Erro:          erro,
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

	retornoExclusao := integracao.Excluir()

	var mensagem string
	var sucesso bool
	var erro bool

	if retornoExclusao == true {
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
