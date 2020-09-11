package controllers

import (
	"net/http"
	"strconv"

	"github.com/tayron/integra-sistema/models"

	appTemplate "github.com/tayron/integra-sistema/bootstrap/library/template"
)

// IndexApplication -
func IndexApplication(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	var integracaoModel models.Integracao
	var numeroIntegracoes int = integracaoModel.ObterNumeroIntegracoes()

	var usuarioModel models.Usuario
	var numeroUSuarios int = usuarioModel.ObterNumeroUsuarios()

	var logModel models.Log
	var numeroIntegracoesRealizadas int = logModel.ObterNumeroIntegracoesRealizadas()

	parametros := appTemplate.Parametro{
		System: appTemplate.ObterInformacaoSistema(w, r),
		Parametro: map[string]string{
			"numeroIntegracoes":           strconv.Itoa(numeroIntegracoes),
			"numeroUsuarios":              strconv.Itoa(numeroUSuarios),
			"numeroIntegracoesRealizadas": strconv.Itoa(numeroIntegracoesRealizadas),
		},
	}

	appTemplate.LoadView(w, "", "homePage", parametros)
}
