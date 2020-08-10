package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("t0op-s3cr3t"))

// SetSessionData - Grava um dado na sessão
func SetSessionData(key string, value string, w http.ResponseWriter, r *http.Request) {
	sessionStore, _ := store.Get(r, "session")
	sessionStore.Values[key] = value
	sessionStore.Save(r, w)
}

// GetSessionData - Retorna um dado da sessão
func GetSessionData(key string, w http.ResponseWriter, r *http.Request) string {
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

// ClearSessionData - Limpa todos os dados da sessão
func ClearSessionData(w http.ResponseWriter, r *http.Request) {
	sessionStore, _ := store.Get(r, "session")
	sessionStore.Options.MaxAge = -1
	sessionStore.Save(r, w)
}
