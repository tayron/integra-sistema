package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/tayron/integra-sistema/models"
)

// EnviarRequisicaoViaPOST - Envia dados para uma api passando dados via POST
func EnviarRequisicaoViaPOST(integracao models.Integracao, metodo string, listaParametros []models.Parametro, r *http.Request) ([]byte, bool) {
	formData := url.Values{}

	for _, parametro := range listaParametros {
		param := []string{r.FormValue(parametro.NomeParametroEntrada)}
		formData[parametro.NomeParametroSaida] = param
	}

	resp, err := http.PostForm(integracao.APISistemaDestino, formData)

	if err != nil {
		var mensagemErro string = err.Error()
		return []byte(mensagemErro), false
	}

	retornoAPI, _ := ioutil.ReadAll(resp.Body)
	data, _ := json.Marshal(formData)
	retornoString := strings.ReplaceAll(string(retornoAPI), `\"`, `"`)

	gravarLogIntegracao(
		integracao.ID,
		integracao.APISistemaDestino,
		metodo,
		fmt.Sprintf("%s", data),
		retornoString,
	)

	return []byte(retornoString), true
}

// EnviarRequisicaoJSONViaPOST - Envia dados para uma api passando dados via JSON via POST
func EnviarRequisicaoJSONViaPOST(integracao models.Integracao, metodo string, listaParametros []models.Parametro, r *http.Request) ([]byte, bool) {
	jsonData := map[string]string{}

	for _, parametro := range listaParametros {
		jsonData[parametro.NomeParametroSaida] = r.FormValue(parametro.NomeParametroEntrada)
	}

	jsonValue, _ := json.Marshal(jsonData)
	response, _ := http.Post(integracao.APISistemaDestino, "application/json", bytes.NewBuffer(jsonValue))
	retornoAPI, _ := ioutil.ReadAll(response.Body)

	data, _ := json.Marshal(jsonData)
	retornoString := strings.ReplaceAll(string(retornoAPI), `\"`, `"`)

	gravarLogIntegracao(
		integracao.ID,
		integracao.APISistemaDestino,
		metodo,
		fmt.Sprintf("%s", data),
		retornoString,
	)

	return retornoAPI, true
}

// EnviarRequisicaoViaGET - Envia dados para uma api passando dados via GET
func EnviarRequisicaoViaGET(integracao models.Integracao, metodo string, listaParametros []models.Parametro, r *http.Request) ([]byte, bool) {
	req, _ := http.NewRequest("GET", integracao.APISistemaDestino, nil)
	query := req.URL.Query()

	for _, parametro := range listaParametros {
		query.Add(parametro.NomeParametroSaida, r.FormValue(parametro.NomeParametroEntrada))
	}

	req.URL.RawQuery = query.Encode()
	resp, _ := http.Get(req.URL.String())

	retornoAPI, _ := ioutil.ReadAll(resp.Body)
	retornoString := strings.ReplaceAll(string(retornoAPI), `\"`, `"`)

	gravarLogIntegracao(integracao.ID, req.URL.String(), metodo, "", retornoString)
	return retornoAPI, true
}

func gravarLogIntegracao(idIntegracao int, apiDestino string, metodo string, parametro string, retornoIntegracao string) {
	log := models.Log{
		IntegracaoID: idIntegracao,
		APIDestino:   apiDestino,
		Metodo:       metodo,
		Parametro:    parametro,
		Resposta:     html.EscapeString(retornoIntegracao),
	}

	log.Gravar()
}
