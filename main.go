package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tayron/integra-sistema/configuration"
	"github.com/tayron/integra-sistema/handlers"
	"github.com/tayron/integra-sistema/models"
)

func init() {
	// Configuração da aplicação
	configuration.SetConfiguracao("nomeAplicativo", "Integra Sistema")
	configuration.SetConfiguracao("versaoAplicativo", "1.0")
	configuration.SetConfiguracao("portaServidorExecucao", os.Getenv("PORTA_SERVIDOR"))

	// Configuração banco de dados
	configuration.SetConfiguracao("localhost", "servidor_mysql_local")
	configuration.SetConfiguracao("porta", "3306")
	configuration.SetConfiguracao("usuario", "root")
	configuration.SetConfiguracao("senha", "yakTLS&70c52")
	configuration.SetConfiguracao("banco", "cursogo")

	if configuration.GetConfiguracao("portaServidorExecucao") == "" {
		panic("Deve-se informar a porta do servidor: PORTA_SERVIDOR=80 go run *.go")
	}

	models.CriarTabelaIntegracao()
}

func main() {
	http.HandleFunc("/", handlers.InicioHandler)

	fmt.Printf("Servidor executando em: http://127.0.0.1:%s\n", configuration.GetConfiguracao("portaServidorExecucao"))
	enderecoServidor := fmt.Sprintf(":%s", configuration.GetConfiguracao("portaServidorExecucao"))

	log.Fatalln(http.ListenAndServe(enderecoServidor, nil))
}
