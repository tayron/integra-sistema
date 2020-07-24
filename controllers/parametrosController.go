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

// ListarParametro -
func ListarParametro(w http.ResponseWriter, r *http.Request) {

	parametrosURL := mux.Vars(r)
	idIntegracao, _ := strconv.ParseInt(parametrosURL["id"], 10, 64)

	integracao := models.Integracao{}
	parametro := models.Parametro{}

	parametros := struct {
		NomeSistema     string
		VersaoSistema   string
		Mensagem        string
		Sucesso         bool
		Erro            bool
		Integracao      models.Integracao
		ListaParametros []models.Parametro
	}{
		NomeSistema:     os.Getenv("NOME_SISTEMA"),
		VersaoSistema:   os.Getenv("VERSAO_SISTEMA"),
		Integracao:      integracao.BuscarPorID(idIntegracao),
		ListaParametros: parametro.BuscarPorIDIntegracao(idIntegracao),
	}

	var templates = template.Must(template.ParseGlob("template/*.html"))
	template.Must(templates.ParseGlob("template/layout/*.html"))
	template.Must(templates.ParseGlob("template/parametro/*.html"))
	templates.ExecuteTemplate(w, "listarParametroPage", parametros)
}

// CriarParametro -
func CriarParametro(w http.ResponseWriter, r *http.Request) {

	integracaoID, _ := strconv.ParseInt(r.FormValue("integracao_id"), 10, 64)
	nomeParametroEntrada := r.FormValue("nome_parametro_entrada")
	nomeParametroSaida := r.FormValue("nome_parametro_saida")

	parametro := models.Parametro{
		IntegracaoID:         integracaoID,
		NomeParametroEntrada: nomeParametroEntrada,
		NomeParametroSaida:   nomeParametroSaida,
	}

	retornoGravacao := parametro.Gravar()

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

	integracao := models.Integracao{}

	parametros := struct {
		NomeSistema     string
		VersaoSistema   string
		Mensagem        string
		Sucesso         bool
		Erro            bool
		Integracao      models.Integracao
		ListaParametros []models.Parametro
	}{
		NomeSistema:     os.Getenv("NOME_SISTEMA"),
		VersaoSistema:   os.Getenv("VERSAO_SISTEMA"),
		Sucesso:         sucesso,
		Erro:            erro,
		Mensagem:        mensagem,
		Integracao:      integracao.BuscarPorID(integracaoID),
		ListaParametros: parametro.BuscarPorIDIntegracao(integracaoID),
	}

	var templates = template.Must(template.ParseGlob("template/*.html"))
	template.Must(templates.ParseGlob("template/layout/*.html"))
	template.Must(templates.ParseGlob("template/parametro/*.html"))
	templates.ExecuteTemplate(w, "listarParametroPage", parametros)
}
