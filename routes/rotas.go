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
	r.HandleFunc("/integracao/editar/{id:[0-9]+}", handlers.EditarIntegracaoHandler).Methods("GET")
	r.HandleFunc("/integracao/editar/{id:[0-9]+}", handlers.GravarIntegracaoHandler).Methods("POST")
	r.HandleFunc("/integracao/excluir", handlers.ExcluirIntegracaoHandler).Methods("POST")
	r.HandleFunc("/parametros/integracao/{id:[0-9]+}", handlers.GerirParametrosHandler).Methods("GET")
	r.HandleFunc("/parametro/cadastrar", handlers.CriarParametroHandler).Methods("POST")
	http.Handle("/", r)
}
