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
}

// BootstrapApplication -
func BootstrapApplication() {
	routes.CarregarRotas()
	StartarServidor()
}
