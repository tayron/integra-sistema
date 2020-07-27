package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tayron/integra-sistema/controllers"
)

// CarregarRotas -
func CarregarRotas() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.IndexApplication).Methods("GET")
	r.HandleFunc("/", controllers.CriarIntegracao).Methods("POST")
	r.HandleFunc("/integracoes", controllers.ListarIntegracao).Methods("GET")
	r.HandleFunc("/integracao/editar/{id:[0-9]+}", controllers.EditarIntegracao).Methods("GET")
	r.HandleFunc("/integracao/editar/{id:[0-9]+}", controllers.SalvarIntegracao).Methods("POST")
	r.HandleFunc("/integracao/excluir", controllers.ExcluirIntegracao).Methods("POST")
	r.HandleFunc("/parametros/integracao/{id:[0-9]+}", controllers.ListarParametro).Methods("GET")
	r.HandleFunc("/parametros/integracao/{id:[0-9]+}", controllers.CriarParametro).Methods("POST")
	r.HandleFunc("/logs/integracao/{id:[0-9]+}", controllers.ListarLog).Methods("GET")
	r.HandleFunc("/api/{endpoint}", controllers.ProcessarIntegracao)
	http.Handle("/", r)
}
