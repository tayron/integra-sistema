package bootstrap

import (
	"github.com/joho/godotenv"
	"github.com/tayron/integra-sistema/models"
	"github.com/tayron/integra-sistema/routes"
)

func init() {
	godotenv.Load()
	models.CriarTabelaIntegracao()
	models.CriarTabelaParametro()
	models.CriarTabelaLog()
	models.CriarTabelaUsuario()
}

// BootstrapApplication -
func BootstrapApplication() {
	routes.CarregarRotas()
	StartarServidor()
}
