package bootstrap

import (
	"github.com/joho/godotenv"
	"github.com/tayron/integra-sistema/bootstrap/library/routes"
	"github.com/tayron/integra-sistema/models"
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
