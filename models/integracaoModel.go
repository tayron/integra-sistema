package models

import (
	"time"

	"github.com/tayron/integra-sistema/database"
)

type Integracao struct {
	ID                   int
	Nome                 string
	NomeSistemaOrigem    string
	APISistemaOrigem     string
	MetodoSistemaOrigem  string
	NomeSistemaDestino   string
	APISistemaDestino    string
	MetodoSistemaDestino string
	Status               bool
	Criacao              time.Time
	Alteracao            time.Time
}

// CriarTabelaIntegracao -
func CriarTabelaIntegracao() {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `create table if not exists integracoes (
		id integer auto_increment,
		nome varchar(255),
		nome_sistema_origem varchar(255),
		api_sistema_origem varchar(255),
		metodo_sistema_origem char(6),
		nome_sistema_destino varchar(255),
		api_sistema_destino varchar(255),
		metodo_sistema_destino char(6),			
		status bool DEFAULT 1,	
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
func (i Integracao) Gravar() bool {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `insert into integracoes 
	(nome, nome_sistema_origem, api_sistema_origem, metodo_sistema_origem, 
	nome_sistema_destino, api_sistema_destino, metodo_sistema_destino) 
	values (?, ?, ?, ?, ?, ?, ?)`

	stmt, _ := db.Prepare(sql)

	resultado, err := stmt.Exec(
		i.Nome,
		i.NomeSistemaOrigem,
		i.APISistemaOrigem,
		i.MetodoSistemaOrigem,
		i.NomeSistemaDestino,
		i.APISistemaDestino,
		i.MetodoSistemaDestino)

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
func (i Integracao) Atualizar() bool {

	db := database.ObterConexao()
	defer db.Close()

	var sql string = `UPDATE integracoes SET
	nome = ?, nome_sistema_origem = ?, api_sistema_origem = ?, metodo_sistema_origem = ?,
	nome_sistema_destino = ?, api_sistema_destino = ?, metodo_sistema_destino = ?
	WHERE id = ?`

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	resultado, err := stmt.Exec(
		i.Nome,
		i.NomeSistemaOrigem,
		i.APISistemaOrigem,
		i.MetodoSistemaOrigem,
		i.NomeSistemaDestino,
		i.APISistemaDestino,
		i.MetodoSistemaDestino,
		i.ID)

	if err != nil {
		panic(err)
	}

	numeroRegistrosAlterados, err := resultado.RowsAffected()

	if err != nil {
		panic(err)
	}

	if numeroRegistrosAlterados > 0 {
		return true
	}

	return true
}

// Excluir -
func (i Integracao) Excluir() bool {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `DELETE FROM integracoes WHERE id = ?`

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	resultado, err := stmt.Exec(i.ID)

	if err != nil {
		panic(err)
	}

	numeroRegistrosAlterados, err := resultado.RowsAffected()

	if err != nil {
		panic(err)
	}

	if numeroRegistrosAlterados > 0 {
		return true
	}

	return true
}

// BuscarTodos -
func (i Integracao) BuscarTodos() []Integracao {

	db := database.ObterConexao()
	defer db.Close()

	var sql string = `SELECT id, nome, nome_sistema_origem, api_sistema_origem, metodo_sistema_origem, 
	nome_sistema_destino, api_sistema_destino, metodo_sistema_destino, status FROM integracoes ORDER BY id DESC`

	rows, _ := db.Query(sql)
	defer rows.Close()

	var listaIntegracoes []Integracao

	for rows.Next() {
		var integracao Integracao

		rows.Scan(&integracao.ID,
			&integracao.Nome,
			&integracao.NomeSistemaOrigem,
			&integracao.APISistemaOrigem,
			&integracao.MetodoSistemaOrigem,
			&integracao.NomeSistemaDestino,
			&integracao.APISistemaDestino,
			&integracao.MetodoSistemaDestino,
			&integracao.Status)

		listaIntegracoes = append(listaIntegracoes, integracao)
	}

	return listaIntegracoes
}

// BuscarPorID -
func (i Integracao) BuscarPorID(idIntegracao int) Integracao {

	db := database.ObterConexao()
	defer db.Close()

	var sql string = `SELECT id, nome, nome_sistema_origem, api_sistema_origem, metodo_sistema_origem, 
	nome_sistema_destino, api_sistema_destino, metodo_sistema_destino, status FROM integracoes WHERE id = ?`

	rows, _ := db.Query(sql, idIntegracao)
	defer rows.Close()

	var integracao Integracao
	for rows.Next() {
		rows.Scan(&integracao.ID,
			&integracao.Nome,
			&integracao.NomeSistemaOrigem,
			&integracao.APISistemaOrigem,
			&integracao.MetodoSistemaOrigem,
			&integracao.NomeSistemaDestino,
			&integracao.APISistemaDestino,
			&integracao.MetodoSistemaDestino,
			&integracao.Status)

		return integracao
	}

	return integracao
}
