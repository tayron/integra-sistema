package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tayron/gopaginacao"
	appTemplate "github.com/tayron/integra-sistema/bootstrap/library/template"
	"github.com/tayron/integra-sistema/models"
)

// ListarIntegracao -
func ListarIntegracao(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	integracaoModel := models.Integracao{}

	numeroTotalRegistro := integracaoModel.ObterNumeroIntegracoes()
	htmlPaginacao, offset, err := gopaginacao.CriarPaginacao(numeroTotalRegistro, r)

	var listaIntegracoes []models.Integracao

	if err == nil {
		listaIntegracoes = integracaoModel.BuscarTodos(offset)
	}

	var Integracoes = struct {
		Integracoes []models.Integracao
		Paginacao   template.HTML
	}{
		Integracoes: listaIntegracoes,
		Paginacao:   template.HTML(htmlPaginacao),
	}

	parametros := appTemplate.Parametro{
		System:    appTemplate.ObterInformacaoSistema(w, r),
		Parametro: Integracoes,
	}

	appTemplate.LoadView(w, "template/integracao/*.html", "listarIntegracoesPage", parametros)
}

// CadastrarIntegracao -
func CadastrarIntegracao(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	flashMessage := appTemplate.FlashMessage{}

	if r.Method == "POST" {
		integracao := models.Integracao{
			Nome:                 r.FormValue("nome"),
			Endpoint:             r.FormValue("endpoint"),
			NomeSistemaOrigem:    r.FormValue("nome_sistema_origem"),
			NomeSistemaDestino:   r.FormValue("nome_sistema_destino"),
			APISistemaDestino:    r.FormValue("api_sistema_destino"),
			MetodoSistemaDestino: r.FormValue("metodo_sistema_destino"),
		}

		retornoGravacao := integracao.Gravar()

		if retornoGravacao == true {
			flashMessage.Type, flashMessage.Message = appTemplate.ObterTipoMensagemGravacaoSucesso()
		} else {
			flashMessage.Type, flashMessage.Message = appTemplate.ObterTipoMensagemGravacaoErro()
		}
	}

	integracaoModel := models.Integracao{}

	numeroTotalRegistro := integracaoModel.ObterNumeroIntegracoes()
	htmlPaginacao, offset, err := gopaginacao.CriarPaginacao(numeroTotalRegistro, r)

	var listaIntegracoes []models.Integracao

	if err == nil {
		listaIntegracoes = integracaoModel.BuscarTodos(offset)
	}

	var Integracoes = struct {
		Integracoes []models.Integracao
		Paginacao   template.HTML
	}{
		Integracoes: listaIntegracoes,
		Paginacao:   template.HTML(htmlPaginacao),
	}

	parametros := appTemplate.Parametro{
		System:       appTemplate.ObterInformacaoSistema(w, r),
		FlashMessage: flashMessage,
		Parametro:    Integracoes,
	}

	appTemplate.LoadView(w, "template/integracao/*.html", "cadastrarIntegracoesPage", parametros)
}

// EditarIntegracao -
func EditarIntegracao(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	parametrosURL := mux.Vars(r)
	idIntegracao, _ := strconv.ParseInt(parametrosURL["id"], 10, 64)
	flashMessage := appTemplate.FlashMessage{}

	if r.Method == "POST" {
		id, _ := strconv.Atoi(r.FormValue("id"))
		integracaoModel := models.Integracao{
			ID:                   id,
			Nome:                 r.FormValue("nome"),
			Endpoint:             r.FormValue("endpoint"),
			NomeSistemaOrigem:    r.FormValue("nome_sistema_origem"),
			NomeSistemaDestino:   r.FormValue("nome_sistema_destino"),
			APISistemaDestino:    r.FormValue("api_sistema_destino"),
			MetodoSistemaDestino: r.FormValue("metodo_sistema_destino"),
		}

		retornoGravacao := integracaoModel.Atualizar()

		if retornoGravacao == true {
			flashMessage.Type, flashMessage.Message = appTemplate.ObterTipoMensagemGravacaoSucesso()
		} else {
			flashMessage.Type, flashMessage.Message = appTemplate.ObterTipoMensagemGravacaoErro()
		}
	}

	integracao := models.Integracao{}

	var Integracao = struct {
		Integracao models.Integracao
	}{
		Integracao: integracao.BuscarPorID(idIntegracao),
	}

	parametros := appTemplate.Parametro{
		System:       appTemplate.ObterInformacaoSistema(w, r),
		FlashMessage: flashMessage,
		Parametro:    Integracao,
	}

	appTemplate.LoadView(w, "template/integracao/*.html", "editarIntegracoesPage", parametros)
}

// ExcluirIntegracao -
func ExcluirIntegracao(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	idIntegracao, _ := strconv.Atoi(r.FormValue("id"))
	flashMessage := appTemplate.FlashMessage{}

	integracaoModel := models.Integracao{
		ID: idIntegracao,
	}

	retornoExclusao := integracaoModel.Excluir()

	if retornoExclusao == true {
		flashMessage.Type, flashMessage.Message = appTemplate.ObterTipoMensagemExclusaoSucesso()
	} else {
		flashMessage.Type, flashMessage.Message = appTemplate.ObterTipoMensagemExclusaoErro()
	}

	numeroTotalRegistro := integracaoModel.ObterNumeroIntegracoes()
	htmlPaginacao, offset, err := gopaginacao.CriarPaginacao(numeroTotalRegistro, r)

	var listaIntegracoes []models.Integracao

	if err == nil {
		listaIntegracoes = integracaoModel.BuscarTodos(offset)
	}

	var Integracoes = struct {
		Integracoes []models.Integracao
		Paginacao   template.HTML
	}{
		Integracoes: listaIntegracoes,
		Paginacao:   template.HTML(htmlPaginacao),
	}

	parametros := appTemplate.Parametro{
		System:       appTemplate.ObterInformacaoSistema(w, r),
		FlashMessage: flashMessage,
		Parametro:    Integracoes,
	}

	appTemplate.LoadView(w, "template/integracao/*.html", "listarIntegracoesPage", parametros)
}
