package models

import (
	"time"

	"github.com/tayron/integra-sistema/bootstrap/library/database"
)

type Integracao struct {
	ID                   int
	Nome                 string
	Endpoint             string
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

// CriarTabelaIntegracao - Caso não existe, cria tabela integração no banco
func CriarTabelaIntegracao() {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `create table if not exists integracoes (
		id integer auto_increment,
		nome varchar(255) NOT NULL UNIQUE,
		endpoint varchar(255) NOT NULL UNIQUE,
		nome_sistema_origem varchar(255) NOT NULL,				
		nome_sistema_destino varchar(255) NOT NULL,
		api_sistema_destino varchar(255) NOT NULL,
		metodo_sistema_destino char(85) NOT NULL,			
		status bool NOT NULL DEFAULT 1,	
		criacao datetime DEFAULT CURRENT_TIMESTAMP,	
		alteracao datetime ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (id)
	)`

	database.ExecutarQuery(db, sql)
}

// Gravar - Grava uma integração
func (i Integracao) Gravar() bool {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `insert into integracoes 
	(nome, endpoint, nome_sistema_origem, nome_sistema_destino, 
	api_sistema_destino, metodo_sistema_destino) 
	values (?, ?, ?, ?, ?, ?)`

	stmt, _ := db.Prepare(sql)

	resultado, err := stmt.Exec(
		i.Nome,
		i.Endpoint,
		i.NomeSistemaOrigem,
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

// Atualizar - Atualiza uma integração
func (i Integracao) Atualizar() bool {

	db := database.ObterConexao()
	defer db.Close()

	var sql string = `UPDATE integracoes SET
	nome = ?, endpoint = ?, nome_sistema_origem = ?,
	nome_sistema_destino = ?, api_sistema_destino = ?, 
	metodo_sistema_destino = ?
	WHERE id = ?`

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	resultado, err := stmt.Exec(
		i.Nome,
		i.Endpoint,
		i.NomeSistemaOrigem,
		i.NomeSistemaDestino,
		i.APISistemaDestino,
		i.MetodoSistemaDestino,
		i.ID)

	if err != nil {
		panic(err)
	}

	_, err = resultado.RowsAffected()

	if err != nil {
		return false
	}

	return true
}

// Excluir - Exclui uma integração
func (i Integracao) Excluir() bool {
	db := database.ObterConexao()
	defer db.Close()

	stmt, _ := db.Prepare("DELETE FROM integracoes WHERE id = ?")
	var _, err = stmt.Exec(i.ID)

	if err != nil {
		return false
	}

	return true
}

// BuscarTodos -Busca todas as integrações
func (i Integracao) BuscarTodos() []Integracao {

	db := database.ObterConexao()
	defer db.Close()

	var sql string = `SELECT id, nome, endpoint, nome_sistema_origem, 
	nome_sistema_destino, api_sistema_destino, metodo_sistema_destino, 
	status FROM integracoes ORDER BY id DESC`

	rows, _ := db.Query(sql)
	defer rows.Close()

	var listaIntegracoes []Integracao

	for rows.Next() {
		var integracao Integracao

		rows.Scan(&integracao.ID,
			&integracao.Nome,
			&integracao.Endpoint,
			&integracao.NomeSistemaOrigem,
			&integracao.NomeSistemaDestino,
			&integracao.APISistemaDestino,
			&integracao.MetodoSistemaDestino,
			&integracao.Status)

		listaIntegracoes = append(listaIntegracoes, integracao)
	}

	return listaIntegracoes
}

// BuscarPorID - Busca integração por id
func (i Integracao) BuscarPorID(idIntegracao int64) Integracao {

	db := database.ObterConexao()
	defer db.Close()

	var sql string = `SELECT id, nome, endpoint, nome_sistema_origem, 
	nome_sistema_destino, api_sistema_destino, metodo_sistema_destino, 
	status FROM integracoes WHERE id = ?`

	rows, _ := db.Query(sql, idIntegracao)
	defer rows.Close()

	var integracao Integracao
	for rows.Next() {
		rows.Scan(&integracao.ID,
			&integracao.Nome,
			&integracao.Endpoint,
			&integracao.NomeSistemaOrigem,
			&integracao.NomeSistemaDestino,
			&integracao.APISistemaDestino,
			&integracao.MetodoSistemaDestino,
			&integracao.Status)

		return integracao
	}

	return integracao
}

// BuscarPorEndpoint - Busca integraçaõ através do endpoint
func (i Integracao) BuscarPorEndpoint(endpoint string) Integracao {

	db := database.ObterConexao()
	defer db.Close()

	var sql string = `SELECT id, nome, endpoint, nome_sistema_origem, 
	nome_sistema_destino, api_sistema_destino, metodo_sistema_destino, 
	status FROM integracoes WHERE endpoint = ?`

	rows, _ := db.Query(sql, endpoint)
	defer rows.Close()

	var integracao Integracao

	for rows.Next() {
		rows.Scan(&integracao.ID,
			&integracao.Nome,
			&integracao.Endpoint,
			&integracao.NomeSistemaOrigem,
			&integracao.NomeSistemaDestino,
			&integracao.APISistemaDestino,
			&integracao.MetodoSistemaDestino,
			&integracao.Status)

		return integracao
	}

	return integracao
}
