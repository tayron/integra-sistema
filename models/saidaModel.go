package models

import (
	"github.com/tayron/integra-sistema/database"
)

type Saida struct {
	ID                         int
	NomeAtributoSistemaDestino string
}

// CriarTabelaSaida -
func CriarTabelaSaida() {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `create table if not exists saidas (
		id integer auto_increment,		
		nome_atributo_sistema_destino varchar(255),
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
func (s Saida) Gravar() int64 {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `insert into saidas 
	(nome_atributo_sistema_destino) 
	values (?)`

	stmt, _ := db.Prepare(sql)

	resultado, err := stmt.Exec(
		s.NomeAtributoSistemaDestino)

	idRegistro, err := resultado.LastInsertId()

	if err != nil {
		panic(err)
	}

	return idRegistro
}

// Excluir -
func (s Saida) Excluir() bool {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `DELETE FROM saidas WHERE id = ?`

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	resultado, err := stmt.Exec(s.ID)

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
