package models

import (
	"github.com/tayron/integra-sistema/database"
)

type Parametro struct {
	ID                   int
	IntegracaoID         int64
	NomeParametroEntrada string
	NomeParametroSaida   string
}

// CriarTabelaParametro -
func CriarTabelaParametro() {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `create table if not exists parametros (
		id integer auto_increment,
		integracao_id integer,
		nome_parametro_entrada varchar(255),
		nome_parametro_saida varchar(255),		
		criacao datetime DEFAULT CURRENT_TIMESTAMP,	
		alteracao datetime ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (id),
		FOREIGN KEY (integracao_id) REFERENCES integracoes(id)
	)`
	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}

// Gravar -
func (p Parametro) Gravar() bool {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `insert into parametros 
		(nome_parametro_entrada, nome_parametro_saida, integracao_id) 
		values (?, ?, ?)`

	stmt, _ := db.Prepare(sql)

	resultado, err := stmt.Exec(
		p.NomeParametroEntrada,
		p.NomeParametroSaida,
		p.IntegracaoID)

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
func (p Parametro) Excluir() bool {
	db := database.ObterConexao()
	defer db.Close()

	stmt, _ := db.Prepare("DELETE FROM parametros WHERE id = ?")
	var _, err = stmt.Exec(p.ID)

	if err != nil {
		return false
	}

	return true

}

// BuscarPorIDIntegracao -
func (p Parametro) BuscarPorIDIntegracao(idIntegracao int64) []Parametro {

	db := database.ObterConexao()
	defer db.Close()

	var sql string = `SELECT id, integracao_id, nome_parametro_entrada, nome_parametro_saida 
		FROM parametros WHERE integracao_id = ? ORDER BY id DESC`

	rows, _ := db.Query(sql, idIntegracao)
	defer rows.Close()

	var listaParametro []Parametro
	for rows.Next() {

		var parametro Parametro

		rows.Scan(&parametro.ID,
			&parametro.IntegracaoID,
			&parametro.NomeParametroEntrada,
			&parametro.NomeParametroSaida)

		listaParametro = append(listaParametro, parametro)
	}

	return listaParametro
}
