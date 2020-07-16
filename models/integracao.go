package models

import (
	"log"

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
		PRIMARY KEY (id)
	)`

	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}

// Gravar -
func (i Integracao) Gravar(integracao Integracao) bool {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `insert into integracoes 
	(nome, nome_sistema_origem, api_sistema_origem, metodo_sistema_origem, 
	nome_sistema_destino, api_sistema_destino, metodo_sistema_destino) 
	values (?, ?, ?, ?, ?, ?, ? )`

	stmt, _ := db.Prepare(sql)

	resultado, err := stmt.Exec(
		integracao.Nome,
		integracao.NomeSistemaOrigem,
		integracao.APISistemaOrigem,
		integracao.MetodoSistemaOrigem,
		integracao.NomeSistemaDestino,
		integracao.APISistemaDestino,
		integracao.MetodoSistemaDestino)

	if err != nil {
		log.Println(err)
	}

	if resultado.LastInsertId != nil {
		return true
	}

	return false
}

// BuscarTodos -
func (i Integracao) BuscarTodos() []Integracao {

	db := database.ObterConexao()
	defer db.Close()

	var sql string = `SELECT id, nome, nome_sistema_origem, api_sistema_origem, metodo_sistema_origem, 
	nome_sistema_destino, api_sistema_destino, metodo_sistema_destino FROM integracoes`

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
			&integracao.MetodoSistemaDestino)

		listaIntegracoes = append(listaIntegracoes, integracao)
	}

	return listaIntegracoes
}
