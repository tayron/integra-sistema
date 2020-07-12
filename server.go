package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tayron/Integrador/handlers"
	"github.com/tayron/Integrador/models"
)

func main() {
	portaServidor := fmt.Sprintf(os.Getenv("PORTA_SERVIDOR"))

	if portaServidor == "" {
		panic("Deve-se informar a porta do servidor: PORTA_SERVIDOR=80 go run *.go")
	}

	http.HandleFunc("/", handlers.InicioHandler)

	models.CriarTabelaIntegracao()

	fmt.Printf("Servidor executando em: http://127.0.0.1:%s\n", portaServidor)
	enderecoServidor := fmt.Sprintf(":%s", portaServidor)
	log.Fatalln(http.ListenAndServe(enderecoServidor, nil))
}
