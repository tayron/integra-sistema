package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tayron/integra-sistema/bootstrap/library/template"
	"github.com/tayron/integra-sistema/bootstrap/library/util"
	"github.com/tayron/integra-sistema/models"
)

// ListarUsuario -
func ListarUsuario(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	usuarioModel := models.Usuario{}

	var Usuarios = struct {
		ListaUsuarios []models.Usuario
	}{
		ListaUsuarios: usuarioModel.BuscarTodos(),
	}

	parametros := template.Parametro{
		System:    template.ObterInformacaoSistema(w, r),
		Parametro: Usuarios,
	}

	template.LoadView(w, "template/usuario/*.html", "listarUsuariosPage", parametros)
}

// CadastrarUsuario -
func CadastrarUsuario(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	flashMessage := template.FlashMessage{}

	if r.Method == "POST" {
		senha := r.FormValue("senha")
		senhaCriptografada, _ := util.GerarHashSenha(senha)

		usuarioEntidade := models.Usuario{
			Nome:  r.FormValue("nome"),
			Login: r.FormValue("login"),
			Senha: senhaCriptografada,
		}

		retornoGravacao := usuarioEntidade.Gravar()

		if retornoGravacao == true {
			flashMessage.Type, flashMessage.Message = template.ObterTipoMensagemGravacaoSucesso()
		} else {
			flashMessage.Type, flashMessage.Message = template.ObterTipoMensagemGravacaoErro()
		}
	}

	parametros := template.Parametro{
		System:       template.ObterInformacaoSistema(w, r),
		FlashMessage: flashMessage,
	}

	template.LoadView(w, "template/usuario/*.html", "cadastrarUsuarioPage", parametros)
}

// EditarUsuario -
func EditarUsuario(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	parametrosURL := mux.Vars(r)
	id, _ := strconv.Atoi(parametrosURL["id"])
	flashMessage := template.FlashMessage{}

	if r.Method == "POST" {

		senha := r.FormValue("senha")
		senhaCriptografada := ""

		if senha != "" {
			senhaCriptografada, _ = util.GerarHashSenha(senha)
		}

		var ativo bool = false
		if r.FormValue("ativo") == "1" {
			ativo = true
		}

		usuarioModel := models.Usuario{
			ID:    id,
			Nome:  r.FormValue("nome"),
			Login: r.FormValue("login"),
			Ativo: ativo,
			Senha: senhaCriptografada,
		}

		retornoGravacao := usuarioModel.Atualizar()

		if retornoGravacao == true {
			flashMessage.Type, flashMessage.Message = template.ObterTipoMensagemGravacaoSucesso()
		} else {
			flashMessage.Type, flashMessage.Message = template.ObterTipoMensagemGravacaoErro()
		}
	}

	usuarioModel := models.Usuario{
		ID: id,
	}

	var Usuario = struct {
		Usuario models.Usuario
	}{
		Usuario: usuarioModel.BuscarPorID(),
	}

	parametros := template.Parametro{
		System:       template.ObterInformacaoSistema(w, r),
		FlashMessage: flashMessage,
		Parametro:    Usuario,
	}

	template.LoadView(w, "template/usuario/*.html", "editarUsuarioPage", parametros)
}

// ExcluirUsuario -
func ExcluirUsuario(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	idUsuario, _ := strconv.Atoi(r.FormValue("id"))
	flashMessage := template.FlashMessage{}

	usuarioModel := models.Usuario{
		ID: idUsuario,
	}

	retornoExclusao := usuarioModel.Excluir()

	if retornoExclusao == true {
		flashMessage.Type, flashMessage.Message = template.ObterTipoMensagemExclusaoSucesso()
	} else {
		flashMessage.Type, flashMessage.Message = template.ObterTipoMensagemExclusaoErro()
	}

	var Usuarios = struct {
		ListaUsuarios []models.Usuario
	}{
		ListaUsuarios: usuarioModel.BuscarTodos(),
	}

	parametros := template.Parametro{
		System:       template.ObterInformacaoSistema(w, r),
		FlashMessage: flashMessage,
		Parametro:    Usuarios,
	}

	template.LoadView(w, "template/usuario/*.html", "listarUsuariosPage", parametros)
}
