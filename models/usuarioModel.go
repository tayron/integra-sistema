package models

import (
	"github.com/tayron/integra-sistema/database"
)

type Usuario struct {
	ID    int
	Nome  string
	Login string
	Senha string
	Ativo string
}

// CriarTabelaUsuario -
func CriarTabelaUsuario() {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `create table if not exists usuarios (
		id integer auto_increment,
		nome varchar(255),
		login varchar(255),		
		senha varchar(255),		
		ativo bool DEFAULT 0,		
		criacao datetime DEFAULT CURRENT_TIMESTAMP,	
		alteracao datetime ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (id)
	)`
	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
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
		var sql string = `UPDATE usuarios SET nome = ?, login = ?, ativo = ?, senha = ? WHERE id = ?`

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

	stmt, _ := db.Prepare("DELETE FROM usuarios WHERE id = ?")
	var _, err = stmt.Exec(u.ID)

	if err != nil {
		return false
	}

	return true
}

// BuscarTodos -
func (u Usuario) BuscarTodos() []Usuario {

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

// BuscarPorID -
func (u Usuario) BuscarPorID() Usuario {

	db := database.ObterConexao()
	defer db.Close()

	var sql string = `SELECT id, nome, login, ativo
		FROM usuarios WHERE id = ? ORDER BY id DESC`

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
