package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tayron/integra-sistema/bootstrap/library/template"
	"github.com/tayron/integra-sistema/models"
)

// ListarParametro -
func ListarParametro(w http.ResponseWriter, r *http.Request) {

	parametrosURL := mux.Vars(r)
	idIntegracao, _ := strconv.ParseInt(parametrosURL["id"], 10, 64)

	var mensagem string
	var sucesso bool
	var erro bool

	if r.Method == "POST" {
		integracaoID, _ := strconv.ParseInt(r.FormValue("integracao_id"), 10, 64)
		nomeParametroEntrada := r.FormValue("nome_parametro_entrada")
		nomeParametroSaida := r.FormValue("nome_parametro_saida")

		parametro := models.Parametro{
			IntegracaoID:         integracaoID,
			NomeParametroEntrada: nomeParametroEntrada,
			NomeParametroSaida:   nomeParametroSaida,
		}

		retornoGravacao := parametro.Gravar()

		if retornoGravacao == true {
			sucesso = true
			mensagem = fmt.Sprint("Sucesso ao gravar dados da integração")
		} else {
			erro = true
			mensagem = fmt.Sprint("Erro ao gravar dados da integração")
		}
	}

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
		Mensagem:        mensagem,
		Sucesso:         sucesso,
		Erro:            erro,
		Integracao:      integracao.BuscarPorID(idIntegracao),
		ListaParametros: parametro.BuscarPorIDIntegracao(idIntegracao),
	}

	template.LoadView(w, "template/parametro/*.html", "listarParametroPage", parametros)
}

// ExcluirParametro -
func ExcluirParametro(w http.ResponseWriter, r *http.Request) {
	idIntegracao, _ := strconv.ParseInt(r.FormValue("id_integracao"), 10, 64)

	id, _ := strconv.Atoi(r.FormValue("id_parametro"))
	parametroModel := models.Parametro{
		ID: id,
	}

	retornoExclusao := parametroModel.Excluir()

	var mensagem string
	var sucesso bool
	var erro bool

	if retornoExclusao == true {
		sucesso = true
		mensagem = fmt.Sprint("Sucesso ao excluir o parâmetro")
	} else {
		erro = true
		mensagem = fmt.Sprint("Erro ao excluir o parâmetro")
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
		Mensagem:        mensagem,
		Sucesso:         sucesso,
		Erro:            erro,
		Integracao:      integracao.BuscarPorID(idIntegracao),
		ListaParametros: parametroModel.BuscarPorIDIntegracao(idIntegracao),
	}

	template.LoadView(w, "template/parametro/*.html", "listarParametroPage", parametros)
}
