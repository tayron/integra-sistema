package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/tayron/integra-sistema/handlers"
	"github.com/tayron/integra-sistema/models"
)

func init() {
	godotenv.Load()
	fmt.Println(os.Getenv("NOME_SISTEMA"))
	models.CriarTabelaIntegracao()
}

func main() {
	http.HandleFunc("/", handlers.InicioHandler)

	fmt.Printf("Servidor executando em: http://127.0.0.1:%s\n", os.Getenv("PORTA_SERVIDOR"))
	enderecoServidor := fmt.Sprintf(":%s", os.Getenv("PORTA_SERVIDOR"))

	if os.Getenv("PORTA_SERVIDOR") == "" {
		panic("Deve-se informar a porta de execução do servidor no arquivo .env: PORTA_SERVIDOR=3003")
	}

	log.Fatalln(http.ListenAndServe(enderecoServidor, nil))
}
