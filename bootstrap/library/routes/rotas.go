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

	r.HandleFunc("/usuarios", controllers.ListarUsuario).Methods("GET")
	r.HandleFunc("/usuario/cadastrar", controllers.CadastrarUsuario)
	r.HandleFunc("/usuario/editar/{id:[0-9]+}", controllers.EditarUsuario)
	r.HandleFunc("/usuario/excluir", controllers.ExcluirUsuario).Methods("POST")

	r.HandleFunc("/login", controllers.Login).Methods("GET")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/logout", controllers.Logout).Methods("GET")

	r.HandleFunc("/logs/integracao/{id:[0-9]+}", controllers.ListarLog).Methods("GET")

	r.HandleFunc("/api/{endpoint}", controllers.ProcessarIntegracao)

	http.Handle("/", r)
}
