package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tayron/integra-sistema/bootstrap/library/template"
	"github.com/tayron/integra-sistema/models"
)

// ListarParametro -
func ListarParametro(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	parametrosURL := mux.Vars(r)
	idIntegracao, _ := strconv.ParseInt(parametrosURL["id"], 10, 64)
	flashMessage := template.FlashMessage{}

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
			flashMessage.Type, flashMessage.Message = template.ObterTipoMensagemGravacaoSucesso()
		} else {
			flashMessage.Type, flashMessage.Message = template.ObterTipoMensagemGravacaoErro()
		}
	}

	integracao := models.Integracao{}
	parametro := models.Parametro{}

	var Parametros = struct {
		Integracao      models.Integracao
		ListaParametros []models.Parametro
	}{
		Integracao:      integracao.BuscarPorID(idIntegracao),
		ListaParametros: parametro.BuscarPorIDIntegracao(idIntegracao),
	}

	parametros := template.Parametro{
		System:    template.ObterInformacaoSistema(w, r),
		Parametro: Parametros,
	}

	template.LoadView(w, "template/parametro/*.html", "listarParametroPage", parametros)
}

// ExcluirParametro -
func ExcluirParametro(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	idIntegracao, _ := strconv.ParseInt(r.FormValue("id_integracao"), 10, 64)
	flashMessage := template.FlashMessage{}

	id, _ := strconv.Atoi(r.FormValue("id_parametro"))
	parametroModel := models.Parametro{
		ID: id,
	}

	retornoExclusao := parametroModel.Excluir()

	if retornoExclusao == true {
		flashMessage.Type, flashMessage.Message = template.ObterTipoMensagemExclusaoSucesso()
	} else {
		flashMessage.Type, flashMessage.Message = template.ObterTipoMensagemExclusaoErro()
	}

	integracao := models.Integracao{}
	parametro := models.Parametro{}

	var Parametros = struct {
		Integracao      models.Integracao
		ListaParametros []models.Parametro
	}{
		Integracao:      integracao.BuscarPorID(idIntegracao),
		ListaParametros: parametro.BuscarPorIDIntegracao(idIntegracao),
	}

	parametros := template.Parametro{
		System:       template.ObterInformacaoSistema(w, r),
		FlashMessage: flashMessage,
		Parametro:    Parametros,
	}

	template.LoadView(w, "template/parametro/*.html", "listarParametroPage", parametros)
}
