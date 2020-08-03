package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tayron/integra-sistema/bootstrap/library/template"
	"github.com/tayron/integra-sistema/models"
	"golang.org/x/crypto/bcrypt"
)

func gerarHashSenha(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func compararSenhaComHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func exemploUso() {
	password := "secret"
	hash, _ := gerarHashSenha(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := compararSenhaComHash(password, hash)
	fmt.Println("Match:   ", match)
}

// ListarUsuario -
func ListarUsuario(w http.ResponseWriter, r *http.Request) {
	usuarioModel := models.Usuario{}

	parametros := struct {
		NomeSistema   string
		VersaoSistema string
		Mensagem      string
		Sucesso       bool
		Erro          bool
		ListaUsuarios []models.Usuario
	}{
		NomeSistema:   os.Getenv("NOME_SISTEMA"),
		VersaoSistema: os.Getenv("VERSAO_SISTEMA"),
		ListaUsuarios: usuarioModel.BuscarTodos(),
	}

	template.LoadView(w, "template/usuario/*.html", "listarUsuariosPage", parametros)
}

// CadastrarUsuario -
func CadastrarUsuario(w http.ResponseWriter, r *http.Request) {
	var mensagem string
	var sucesso bool
	var erro bool

	if r.Method == "POST" {
		senha := r.FormValue("senha")
		senhaCriptografada, _ := gerarHashSenha(senha)

		usuarioEntidade := models.Usuario{
			Nome:  r.FormValue("nome"),
			Login: r.FormValue("login"),
			Senha: senhaCriptografada,
		}

		retornoGravacao := usuarioEntidade.Gravar()

		if retornoGravacao == true {
			sucesso = true
			mensagem = fmt.Sprint("Sucesso ao gravar dados do usuário")
		} else {
			erro = true
			mensagem = fmt.Sprint("Erro ao gravar dados do usuário")
		}
	}

	parametros := struct {
		NomeSistema   string
		VersaoSistema string
		Mensagem      string
		Sucesso       bool
		Erro          bool
	}{
		NomeSistema:   os.Getenv("NOME_SISTEMA"),
		VersaoSistema: os.Getenv("VERSAO_SISTEMA"),
		Mensagem:      mensagem,
		Sucesso:       sucesso,
		Erro:          erro,
	}

	template.LoadView(w, "template/usuario/*.html", "cadastrarUsuarioPage", parametros)
}

// EditarUsuario -
func EditarUsuario(w http.ResponseWriter, r *http.Request) {
	parametrosURL := mux.Vars(r)
	id, _ := strconv.Atoi(parametrosURL["id"])

	var mensagem string
	var sucesso bool
	var erro bool

	if r.Method == "POST" {

		senha := r.FormValue("senha")
		senhaCriptografada := ""

		if senha != "" {
			senhaCriptografada, _ = gerarHashSenha(senha)
		}

		usuarioModel := models.Usuario{
			ID:    id,
			Nome:  r.FormValue("nome"),
			Login: r.FormValue("login"),
			Ativo: r.FormValue("ativo"),
			Senha: senhaCriptografada,
		}

		retornoGravacao := usuarioModel.Atualizar()

		if retornoGravacao == true {
			sucesso = true
			mensagem = fmt.Sprint("Sucesso ao gravar dados do usuário")
		} else {
			erro = true
			mensagem = fmt.Sprint("Erro ao gravar dados da usuário")
		}
	}

	usuarioModel := models.Usuario{
		ID: id,
	}

	parametros := struct {
		NomeSistema   string
		VersaoSistema string
		Mensagem      string
		Sucesso       bool
		Erro          bool
		Usuario       models.Usuario
	}{
		NomeSistema:   os.Getenv("NOME_SISTEMA"),
		VersaoSistema: os.Getenv("VERSAO_SISTEMA"),
		Mensagem:      mensagem,
		Sucesso:       sucesso,
		Erro:          erro,
		Usuario:       usuarioModel.BuscarPorID(),
	}

	template.LoadView(w, "template/usuario/*.html", "editarUsuarioPage", parametros)
}

// ExcluirUsuario -
func ExcluirUsuario(w http.ResponseWriter, r *http.Request) {
	idUsuario, _ := strconv.Atoi(r.FormValue("id"))
	var mensagem string
	var sucesso bool
	var erro bool

	usuarioModel := models.Usuario{
		ID: idUsuario,
	}

	retornoExclusao := usuarioModel.Excluir()

	if retornoExclusao == true {
		sucesso = true
		mensagem = fmt.Sprint("Sucesso ao excluir o usuário")
	} else {
		erro = true
		mensagem = fmt.Sprint("Erro ao excluir o usuário")
	}

	parametros := struct {
		NomeSistema   string
		VersaoSistema string
		Mensagem      string
		Sucesso       bool
		Erro          bool
		ListaUsuarios []models.Usuario
	}{
		NomeSistema:   os.Getenv("NOME_SISTEMA"),
		VersaoSistema: os.Getenv("VERSAO_SISTEMA"),
		Mensagem:      mensagem,
		Sucesso:       sucesso,
		Erro:          erro,
		ListaUsuarios: usuarioModel.BuscarTodos(),
	}

	template.LoadView(w, "template/usuario/*.html", "listarUsuariosPage", parametros)
}
