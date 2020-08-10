package controllers

import (
	"net/http"

	"github.com/tayron/integra-sistema/bootstrap/library/session"
	"github.com/tayron/integra-sistema/bootstrap/library/template"
)

// ValidarSessao - Verifica se tem usuário logado e redireciona para tela de login
func ValidarSessao(w http.ResponseWriter, r *http.Request) {
	usuario := session.GetSessionData("login", w, r)

	if usuario == "" {
		http.Redirect(w, r, "/login", 302)
	}
}

// Login - Exibe tela de login e faz loogin do usuário no banco
func Login(w http.ResponseWriter, r *http.Request) {
	flashMessage := template.FlashMessage{}

	if r.Method == "POST" {
		r.ParseForm()
		username := r.PostForm.Get("login")
		password := r.PostForm.Get("senha")

		session.SetSessionData("login", username, w, r)
		session.SetSessionData("senha", password, w, r)

		http.Redirect(w, r, "/", 302)
	}

	parametros := template.Parametro{
		System:       template.ObterInformacaoSistema(),
		FlashMessage: flashMessage,
	}

	template.LoadView(w, "template/autenticacao/*.html", "loginPage", parametros)
}
