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

// GerirParametrosHandler -
func GerirParametrosHandler(w http.ResponseWriter, r *http.Request) {

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
	template.Must(templates.ParseGlob("template/parametro/*.html"))
	err := templates.ExecuteTemplate(w, "listarParametroPage", parametros)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// CriarParametroHandler -
func CriarParametroHandler(w http.ResponseWriter, r *http.Request) {

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
		NomeSistema   string
		VersaoSistema string
		Mensagem      string
		Sucesso       bool
		Erro          bool
		Integracao    models.Integracao
	}{
		NomeSistema:   os.Getenv("NOME_SISTEMA"),
		VersaoSistema: os.Getenv("VERSAO_SISTEMA"),
		Sucesso:       sucesso,
		Erro:          erro,
		Mensagem:      mensagem,
		Integracao:    integracao.BuscarPorID(integracaoID),
	}

	var templates = template.Must(template.ParseGlob("template/*.html"))
	template.Must(templates.ParseGlob("template/layout/*.html"))
	template.Must(templates.ParseGlob("template/parametro/*.html"))
	err := templates.ExecuteTemplate(w, "listarParametroPage", parametros)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
