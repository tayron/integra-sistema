package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tayron/integra-sistema/handlers"
)

// CarregarRotas -
func CarregarRotas() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.InicioHandler).Methods("GET")
	r.HandleFunc("/", handlers.CriarIntegracaoHandler).Methods("POST")
	r.HandleFunc("/integracoes", handlers.ListarIntegracoesHandler).Methods("GET")
	r.HandleFunc("/editar/{id:[0-9]+}", handlers.EditarIntegracaoHandler).Methods("GET")
	r.HandleFunc("/editar/{id:[0-9]+}", handlers.GravarIntegracaoHandler).Methods("POST")
	r.HandleFunc("/excluir-integracao", handlers.ExcluirIntegracaoHandler).Methods("POST")
	r.HandleFunc("/editar-atributos/{id:[0-9]+}", handlers.EditarAtributosHandler).Methods("GET")
	r.HandleFunc("/cadastrar-atributos", handlers.CriarAtributoHandler).Methods("POST")
	http.Handle("/", r)
}
