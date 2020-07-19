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
	http.Handle("/", r)
}
