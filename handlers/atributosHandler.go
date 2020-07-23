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

// EditarAtributosHandler -
func EditarAtributosHandler(w http.ResponseWriter, r *http.Request) {

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
	template.Must(templates.ParseGlob("template/atributo/*.html"))
	err := templates.ExecuteTemplate(w, "editarAtributosPage", parametros)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// CriarAtributoHandler -
func CriarAtributoHandler(w http.ResponseWriter, r *http.Request) {

	integracaoID, _ := strconv.ParseInt(r.FormValue("integracao_id"), 10, 64)
	nomeAtributoSistemaOrigem := r.FormValue("nome_sistema_origem")
	nomeAtributoSistemaDestino := r.FormValue("nome_sistema_destino")

	saida := models.Saida{
		NomeAtributoSistemaDestino: nomeAtributoSistemaDestino,
	}

	saidaID := saida.Gravar()

	entrada := models.Entrada{
		IntegracaoID:              integracaoID,
		SaidaID:                   saidaID,
		NomeAtributoSistemaOrigem: nomeAtributoSistemaOrigem,
	}

	retornoGravacao := entrada.Gravar()

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
	template.Must(templates.ParseGlob("template/atributo/*.html"))
	err := templates.ExecuteTemplate(w, "editarAtributosPage", parametros)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
