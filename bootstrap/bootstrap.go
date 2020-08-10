package bootstrap

import (
	"github.com/joho/godotenv"
	"github.com/tayron/integra-sistema/bootstrap/library/routes"
	"github.com/tayron/integra-sistema/models"
)

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
	routes.CarregarRotas()
	StartarServidor()
}
