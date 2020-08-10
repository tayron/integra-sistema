package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("t0op-s3cr3t"))

// SetDadoSessao - Grava um dado na sessão
func SetDadoSessao(key string, value string, w http.ResponseWriter, r *http.Request) {
	sessionStore, _ := store.Get(r, "session")
	sessionStore.Values[key] = value
	sessionStore.Save(r, w)
}

// GetDadoSessao - Retorna um dado da sessão
func GetDadoSessao(key string, w http.ResponseWriter, r *http.Request) string {
	sessionStore, err := store.Get(r, "session")

	if err != nil {
		return ""
	}

	dado, err2 := sessionStore.Values[key]

	if err2 == false {
		return ""
	}

	return dado.(string)
}

// ClearDadosSessao - Limpa todos os dados da sessão
func ClearDadosSessao(w http.ResponseWriter, r *http.Request) {
	sessionStore, _ := store.Get(r, "session")
	sessionStore.Options.MaxAge = -1
	sessionStore.Save(r, w)
}
