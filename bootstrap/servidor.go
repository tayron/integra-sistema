package bootstrap

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// StartarServidor -
func StartarServidor() {

	fmt.Printf("####################### %s - versão %s #######################\n",
		os.Getenv("NOME_SISTEMA"), os.Getenv("VERSAO_SISTEMA"))

	fmt.Printf("Servidor executando em: http://127.0.0.1:%s\n", os.Getenv("PORTA_SERVIDOR"))

	if os.Getenv("PORTA_SERVIDOR") == "" {
		panic("Deve-se informar a porta de execução do servidor no arquivo .env: PORTA_SERVIDOR=3003")
	}

	enderecoServidor := fmt.Sprintf(":%s", os.Getenv("PORTA_SERVIDOR"))
	log.Fatalln(http.ListenAndServe(enderecoServidor, nil))
}
