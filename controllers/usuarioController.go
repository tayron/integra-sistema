package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tayron/gopaginacao"
	appTemplate "github.com/tayron/integra-sistema/bootstrap/library/template"
	"github.com/tayron/integra-sistema/bootstrap/library/util"
	"github.com/tayron/integra-sistema/models"
)

// ListarUsuario -
func ListarUsuario(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	usuarioModel := models.Usuario{}
	numeroTotalRegistro := usuarioModel.ObterNumeroUsuarios()
	htmlPaginacao, offset, err := gopaginacao.CriarPaginacao(numeroTotalRegistro, r)

	var listaUsuarios []models.Usuario

	if err == nil {
		listaUsuarios = usuarioModel.BuscarTodos(offset)
	}

	var Usuarios = struct {
		ListaUsuarios []models.Usuario
		Paginacao     template.HTML
	}{
		ListaUsuarios: listaUsuarios,
		Paginacao:     template.HTML(htmlPaginacao),
	}

	parametros := appTemplate.Parametro{
		System:    appTemplate.ObterInformacaoSistema(w, r),
		Parametro: Usuarios,
	}

	appTemplate.LoadView(w, "template/usuario/*.html", "listarUsuariosPage", parametros)
}

// CadastrarUsuario -
func CadastrarUsuario(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	flashMessage := appTemplate.FlashMessage{}

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
			flashMessage.Type, flashMessage.Message = appTemplate.ObterTipoMensagemGravacaoSucesso()
		} else {
			flashMessage.Type, flashMessage.Message = appTemplate.ObterTipoMensagemGravacaoErro()
		}
	}

	parametros := appTemplate.Parametro{
		System:       appTemplate.ObterInformacaoSistema(w, r),
		FlashMessage: flashMessage,
	}

	appTemplate.LoadView(w, "template/usuario/*.html", "cadastrarUsuarioPage", parametros)
}

// EditarUsuario -
func EditarUsuario(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	parametrosURL := mux.Vars(r)
	id, _ := strconv.Atoi(parametrosURL["id"])
	flashMessage := appTemplate.FlashMessage{}

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
			flashMessage.Type, flashMessage.Message = appTemplate.ObterTipoMensagemGravacaoSucesso()
		} else {
			flashMessage.Type, flashMessage.Message = appTemplate.ObterTipoMensagemGravacaoErro()
		}
	}

	usuarioModel := models.Usuario{
		ID: id,
	}

	usuario := usuarioModel.BuscarPorID()

	if usuario.ID == 0 {
		http.Redirect(w, r, "/", 302)
	}

	var Usuario = struct {
		Usuario models.Usuario
	}{
		Usuario: usuarioModel.BuscarPorID(),
	}

	parametros := appTemplate.Parametro{
		System:       appTemplate.ObterInformacaoSistema(w, r),
		FlashMessage: flashMessage,
		Parametro:    Usuario,
	}

	appTemplate.LoadView(w, "template/usuario/*.html", "editarUsuarioPage", parametros)
}

// ExcluirUsuario -
func ExcluirUsuario(w http.ResponseWriter, r *http.Request) {
	ValidarSessao(w, r)

	idUsuario, _ := strconv.Atoi(r.FormValue("id"))
	flashMessage := appTemplate.FlashMessage{}

	usuarioModel := models.Usuario{
		ID: idUsuario,
	}

	retornoExclusao := usuarioModel.Excluir()

	if retornoExclusao == true {
		flashMessage.Type, flashMessage.Message = appTemplate.ObterTipoMensagemExclusaoSucesso()
	} else {
		flashMessage.Type, flashMessage.Message = appTemplate.ObterTipoMensagemExclusaoErro()
	}

	numeroTotalRegistro := usuarioModel.ObterNumeroUsuarios()
	htmlPaginacao, offset, err := gopaginacao.CriarPaginacao(numeroTotalRegistro, r)

	var listaUsuarios []models.Usuario

	if err == nil {
		listaUsuarios = usuarioModel.BuscarTodos(offset)
	}

	var Usuarios = struct {
		ListaUsuarios []models.Usuario
		Paginacao     template.HTML
	}{
		ListaUsuarios: listaUsuarios,
		Paginacao: template.HTML(htmlPaginacao),
	}

	parametros := appTemplate.Parametro{
		System:       appTemplate.ObterInformacaoSistema(w, r),
		FlashMessage: flashMessage,
		Parametro:    Usuarios,
	}

	appTemplate.LoadView(w, "template/usuario/*.html", "listarUsuariosPage", parametros)
}
