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

func main() {
	configuracao := configuration.Configuracao{
		NomeAplicativo:     "Integra Sistema",
		VersaoAplicativo:   "1.0",
		IPServidor:         "127.0.0.1",
		PortaServidor:      os.Getenv("PORTA_SERVIDOR"),
		EnderecoBancoDados: os.Getenv("SERVIDOR_BANCO"),
		PortaBancoDados:    os.Getenv("PORTA_BANCO"),
		NomeBancoDados:     os.Getenv("NOME_SERVIDOR"),
		UsuarioBancoDados:  os.Getenv("USUARIO_SERVIDOR"),
		SenhaBancoDados:    os.Getenv("SENHA_SERVIDOR"),
	}

	if configuracao.GetPortaServidor() == "" {
		panic("Deve-se informar a porta do servidor: PORTA_SERVIDOR=80 go run *.go")
	}

	http.HandleFunc("/", handlers.InicioHandler)

	models.CriarTabelaIntegracao()

	fmt.Printf("Servidor executando em: http://%s:%s\n", configuracao.GetIPServidor(), configuracao.GetPortaServidor())
	enderecoServidor := fmt.Sprintf("%s:%s", configuracao.GetIPServidor(), configuracao.GetPortaServidor())

	log.Fatalln(http.ListenAndServe(enderecoServidor, nil))
}
