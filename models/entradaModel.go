package models

import (
	"github.com/tayron/integra-sistema/database"
)

type Entrada struct {
	ID                        int
	Integracao                Integracao
	Saida                     Saida
	NomeAtributoSistemaOrigem string
	NomeCampoOrigem           string
}

// CriarTabelaEntrada -
func CriarTabelaEntrada() {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `create table if not exists entradas (
		id integer auto_increment,
		integracao_id integer,
		saida_id integer,
		nome_atributo_sistema_origem varchar(255),
		criacao datetime DEFAULT CURRENT_TIMESTAMP,	
		alteracao datetime ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (id),
		FOREIGN KEY (integracao_id) REFERENCES integracoes(id),
		FOREIGN KEY (saida_id) REFERENCES saidas(id)
	)`

	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}

// Gravar -
func (e Entrada) Gravar() bool {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `insert into entradas 
	(nome_atributo_sistema_origem, integracao_id, saida_id) 
	values (?, ?, ?)`

	stmt, _ := db.Prepare(sql)

	resultado, err := stmt.Exec(
		e.NomeAtributoSistemaOrigem,
		e.Integracao.ID,
		e.Saida.ID)

	numeroRegistrosAlterados, err := resultado.RowsAffected()

	if err != nil {
		panic(err)
	}

	if numeroRegistrosAlterados > 0 {
		return true
	}

	return false
}

// Excluir -
func (e Entrada) Excluir() bool {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `DELETE FROM entradas WHERE id = ?`

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	resultado, err := stmt.Exec(e.ID)

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
