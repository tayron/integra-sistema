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
	r.HandleFunc("/integracoes", handlers.ListarHandler).Methods("GET")
	r.HandleFunc("/editar/{id:[0-9]+}", handlers.EditarHandler).Methods("GET")
	http.Handle("/", r)
}
