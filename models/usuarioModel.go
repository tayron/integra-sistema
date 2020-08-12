package models

import (
	"github.com/tayron/integra-sistema/bootstrap/library/database"
)

type Usuario struct {
	ID                int
	Nome              string
	Login             string
	Senha             string
	Ativo             bool
	PermiteExclusao   bool
	permiteEdicao     bool
	PermiteSerListado bool
}

// CriarTabelaUsuario - Cria caso não exista tabela usuaários no banco
func CriarTabelaUsuario() {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `create table if not exists usuarios (
		id integer auto_increment,
		nome varchar(255) NOT NULL,
		login varchar(255) NOT NULL UNIQUE,		
		senha varchar(255),		
		ativo bool DEFAULT 0,
		permite_exclusao bool DEFAULT 1,
		permite_edicao bool DEFAULT 1,
		permite_ser_listado bool DEFAULT 1,
		criacao datetime DEFAULT CURRENT_TIMESTAMP,	
		alteracao datetime ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (id)
	)`

	database.ExecutarQuery(db, sql)
}

// CriarUsuarioAdministrador - Cria usuário default do sistema
func CriarUsuarioAdministrador() {
	var usuarioModel Usuario

	listaUsuarios := usuarioModel.BuscarTodosIndependenteStatus()

	if len(listaUsuarios) == 0 {
		db := database.ObterConexao()
		defer db.Close()

		var sql string = `insert into usuarios 
			(nome, login, senha, ativo, permite_exclusao, permite_edicao, permite_ser_listado) 
			values (?, ?, ?, ?, ?, ?, ?)`

		stmt, _ := db.Prepare(sql)

		usuarioModel := Usuario{
			Nome:              "Integra Sistema",
			Login:             "integra-sistema",
			Senha:             "$2a$14$ZN3eWRZs30egm9pwDOucVeBBu28LMoou4JCTf0EsU2pzLCLyshYnu",
			Ativo:             true,
			PermiteExclusao:   false,
			permiteEdicao:     false,
			PermiteSerListado: false,
		}

		_, err := stmt.Exec(
			usuarioModel.Nome,
			usuarioModel.Login,
			usuarioModel.Senha,
			usuarioModel.Ativo,
			usuarioModel.PermiteExclusao,
			usuarioModel.permiteEdicao,
			usuarioModel.PermiteSerListado,
		)

		if err != nil {
			panic(err)
		}
	}
}

// Gravar -
func (u Usuario) Gravar() bool {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `insert into usuarios 
		(nome, login, senha) values (?, ?, ?)`

	stmt, _ := db.Prepare(sql)

	resultado, err := stmt.Exec(
		u.Nome,
		u.Login,
		u.Senha)

	numeroRegistrosAlterados, err := resultado.RowsAffected()

	if err != nil {
		panic(err)
	}

	if numeroRegistrosAlterados > 0 {
		return true
	}

	return false
}

// Atualizar -
func (u Usuario) Atualizar() bool {

	db := database.ObterConexao()
	defer db.Close()

	if u.Senha != "" {
		var sql string = `UPDATE usuarios SET nome = ?, login = ?, ativo = ?, senha = ? WHERE permite_edicao = true AND id = ?`

		stmt, err := db.Prepare(sql)

		if err != nil {
			panic(err)
		}

		resultado, err := stmt.Exec(
			u.Nome,
			u.Login,
			u.Ativo,
			u.Senha,
			u.ID)

		if err != nil {
			panic(err)
		}

		_, err = resultado.RowsAffected()

		if err != nil {
			return false
		}

	} else {
		var sql string = `UPDATE usuarios SET nome = ?, login = ?, ativo = ? WHERE id = ?`

		stmt, err := db.Prepare(sql)

		if err != nil {
			panic(err)
		}

		resultado, err := stmt.Exec(
			u.Nome,
			u.Login,
			u.Ativo,
			u.ID)

		if err != nil {
			panic(err)
		}

		_, err = resultado.RowsAffected()

		if err != nil {
			return false
		}
	}

	return true
}

// Excluir -
func (u Usuario) Excluir() bool {
	db := database.ObterConexao()
	defer db.Close()

	stmt, _ := db.Prepare("DELETE FROM usuarios WHERE permite_exclusao = true AND id = ?")
	var _, err = stmt.Exec(u.ID)

	if err != nil {
		return false
	}

	return true
}

// BuscarTodos -Retorna todos os usuários
func (u Usuario) BuscarTodos() []Usuario {

	db := database.ObterConexao()
	defer db.Close()

	var sql string = `SELECT id, nome, login, ativo
		FROM usuarios WHERE permite_ser_listado = true ORDER BY id DESC`

	rows, _ := db.Query(sql)
	defer rows.Close()

	var listaUsuarios []Usuario
	for rows.Next() {

		var usuarioModel Usuario

		rows.Scan(&usuarioModel.ID,
			&usuarioModel.Nome,
			&usuarioModel.Login,
			&usuarioModel.Ativo)

		listaUsuarios = append(listaUsuarios, usuarioModel)
	}

	return listaUsuarios
}

// BuscarTodosIndependenteStatus -Retorna todos os usuários independente de status
func (u Usuario) BuscarTodosIndependenteStatus() []Usuario {

	db := database.ObterConexao()
	defer db.Close()

	var sql string = `SELECT id, nome, login, ativo
		FROM usuarios ORDER BY id DESC`

	rows, _ := db.Query(sql)
	defer rows.Close()

	var listaUsuarios []Usuario
	for rows.Next() {

		var usuarioModel Usuario

		rows.Scan(&usuarioModel.ID,
			&usuarioModel.Nome,
			&usuarioModel.Login,
			&usuarioModel.Ativo)

		listaUsuarios = append(listaUsuarios, usuarioModel)
	}

	return listaUsuarios
}

// BuscarPorID - Busca usuário por ID
func (u Usuario) BuscarPorID() Usuario {

	db := database.ObterConexao()
	defer db.Close()

	var sql string = `SELECT id, nome, login, ativo
		FROM usuarios WHERE permite_ser_listado = true AND id = ? ORDER BY id DESC`

	rows, _ := db.Query(sql, u.ID)
	defer rows.Close()

	var usuarioModel Usuario
	for rows.Next() {
		rows.Scan(&usuarioModel.ID,
			&usuarioModel.Nome,
			&usuarioModel.Login,
			&usuarioModel.Ativo)
		return usuarioModel
	}

	return usuarioModel
}

// BuscarPorLoginStatus - Busca usuario por login e status
func (u Usuario) BuscarPorLoginStatus() Usuario {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `SELECT id, nome, login, senha, ativo
		FROM usuarios WHERE login = ? AND ativo = ?`

	rows, _ := db.Query(sql, u.Login, u.Ativo)
	defer rows.Close()

	var usuarioModel Usuario
	for rows.Next() {
		rows.Scan(&usuarioModel.ID,
			&usuarioModel.Nome,
			&usuarioModel.Login,
			&usuarioModel.Senha,
			&usuarioModel.Ativo)
		return usuarioModel
	}

	return usuarioModel
}
