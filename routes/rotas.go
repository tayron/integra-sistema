package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tayron/integra-sistema/controllers"
)

// CarregarRotas -
func CarregarRotas() {
	r := mux.NewRouter()

	s := http.StripPrefix("/public/", http.FileServer(http.Dir("./public/")))
	r.PathPrefix("/public/").Handler(s)

	r.HandleFunc("/", controllers.IndexApplication).Methods("GET")

	r.HandleFunc("/integracoes", controllers.ListarIntegracao).Methods("GET")
	r.HandleFunc("/integracao/cadastrar", controllers.CadastrarIntegracao)
	r.HandleFunc("/integracao/editar/{id:[0-9]+}", controllers.EditarIntegracao)
	r.HandleFunc("/integracao/excluir", controllers.ExcluirIntegracao).Methods("POST")

	r.HandleFunc("/parametros/integracao/{id:[0-9]+}", controllers.ListarParametro)
	r.HandleFunc("/parametro/excluir", controllers.ExcluirParametro).Methods("POST")

	r.HandleFunc("/logs/integracao/{id:[0-9]+}", controllers.ListarLog).Methods("GET")

	r.HandleFunc("/api/{endpoint}", controllers.ProcessarIntegracao)

	http.Handle("/", r)
}
