package bootstrap

import (
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/tayron/integra-sistema/models"
)

var store = sessions.NewCookieStore([]byte("1n+3gra-s1s+3ma"))

// Inicializa as configurações básica do sistema
func init() {
	godotenv.Load()
	models.CriarTabelaIntegracao()
	models.CriarTabelaParametro()
	models.CriarTabelaLog()
	models.CriarTabelaUsuario()
}

// StartApplication - Carrega as rotas e inializa a aplicação
func StartApplication() {
	CarregarRotas()
	StartarServidor()
}
