package main

import (
	"github.com/joho/godotenv"
	"github.com/tayron/integra-sistema/models"
	"github.com/tayron/integra-sistema/routes"
	"github.com/tayron/integra-sistema/server"
)

func init() {
	godotenv.Load()
	models.CriarTabelaIntegracao()
}

func main() {
	routes.CarregarRotas()
	server.StartarServidor()
}
