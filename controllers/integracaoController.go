package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tayron/integra-sistema/bootstrap/library/template"
	"github.com/tayron/integra-sistema/models"
)

// ListarIntegracao -
func ListarIntegracao(w http.ResponseWriter, r *http.Request) {

	integracao := models.Integracao{}

	var Integracoes = struct {
		Integracoes []models.Integracao
	}{
		Integracoes: integracao.BuscarTodos(),
	}

	parametros := template.Parametro{
		System:    template.ObterInformacaoSistema(),
		Parametro: Integracoes,
	}

	template.LoadView(w, "template/integracao/*.html", "listarIntegracoesPage", parametros)
}

// CadastrarIntegracao -
func CadastrarIntegracao(w http.ResponseWriter, r *http.Request) {

	flashMessage := template.FlashMessage{}

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
			flashMessage.Type, flashMessage.Message = template.ObterTipoMensagemGravacaoSucesso()
		} else {
			flashMessage.Type, flashMessage.Message = template.ObterTipoMensagemGravacaoErro()
		}
	}

	integracao := models.Integracao{}

	var Integracoes = struct {
		Integracoes []models.Integracao
	}{
		Integracoes: integracao.BuscarTodos(),
	}

	parametros := template.Parametro{
		System:       template.ObterInformacaoSistema(),
		FlashMessage: flashMessage,
		Parametro:    Integracoes,
	}

	template.LoadView(w, "template/integracao/*.html", "cadastrarIntegracoesPage", parametros)
}

// EditarIntegracao -
func EditarIntegracao(w http.ResponseWriter, r *http.Request) {

	parametrosURL := mux.Vars(r)
	idIntegracao, _ := strconv.ParseInt(parametrosURL["id"], 10, 64)
	flashMessage := template.FlashMessage{}

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
			flashMessage.Type, flashMessage.Message = template.ObterTipoMensagemGravacaoSucesso()
		} else {
			flashMessage.Type, flashMessage.Message = template.ObterTipoMensagemGravacaoErro()
		}
	}

	integracao := models.Integracao{}

	var Integracao = struct {
		Integracao models.Integracao
	}{
		Integracao: integracao.BuscarPorID(idIntegracao),
	}

	parametros := template.Parametro{
		System:       template.ObterInformacaoSistema(),
		FlashMessage: flashMessage,
		Parametro:    Integracao,
	}

	template.LoadView(w, "template/integracao/*.html", "editarIntegracoesPage", parametros)
}

// ExcluirIntegracao -
func ExcluirIntegracao(w http.ResponseWriter, r *http.Request) {
	idIntegracao, _ := strconv.Atoi(r.FormValue("id"))
	flashMessage := template.FlashMessage{}

	integracao := models.Integracao{
		ID: idIntegracao,
	}

	retornoExclusao := integracao.Excluir()

	if retornoExclusao == true {
		flashMessage.Type, flashMessage.Message = template.ObterTipoMensagemExclusaoSucesso()
	} else {
		flashMessage.Type, flashMessage.Message = template.ObterTipoMensagemExclusaoErro()
	}

	var Integracoes = struct {
		Integracoes []models.Integracao
	}{
		Integracoes: integracao.BuscarTodos(),
	}

	parametros := template.Parametro{
		System:       template.ObterInformacaoSistema(),
		FlashMessage: flashMessage,
		Parametro:    Integracoes,
	}

	template.LoadView(w, "template/integracao/*.html", "listarIntegracoesPage", parametros)
}
