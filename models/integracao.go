package models

import (
	"log"

	"github.com/tayron/Integrador/database"
)

type Integracao struct {
	id                   int
	nome                 string
	nomeSistemaOrigem    string
	apiSistemaOrigem     string
	metodoSistemaOrigem  string
	nomeSistemaDestino   string
	apiSistemaDestino    string
	metodoSistemaDestino string
}

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

func (i Integracao) Gravar(integracao Integracao) bool {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `insert into integracoes 
	(nome, nome_sistema_origem, api_sistema_origem, metodo_sistema_origem, 
	nome_sistema_destino, api_sistema_destino, metodo_sistema_destino) 
	values (?, ?, ?, ?, ?, ?, ? )`

	stmt, _ := db.Prepare(sql)

	resultado, err := stmt.Exec(integracao.nome, integracao.nomeSistemaOrigem, integracao.apiSistemaOrigem, integracao.metodoSistemaOrigem, integracao.nomeSistemaDestino, integracao.apiSistemaDestino, integracao.metodoSistemaDestino)

	if err != nil {
		log.Println(err)
	}

	if resultado.LastInsertId != nil {
		return true
	}

	return false
}
