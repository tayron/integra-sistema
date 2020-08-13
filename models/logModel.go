package models

import (
	"github.com/tayron/integra-sistema/bootstrap/library/database"
)

type Log struct {
	ID           int
	IntegracaoID int
	APIDestino   string
	Metodo       string
	Parametro    string
	Resposta     string
}

// CriarTabelaLog -
func CriarTabelaLog() {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `create table if not exists logs (
		id integer auto_increment,
		integracao_id integer,
		api_destino varchar(255),
		metodo varchar(255),
		parametro text,		
		resposta text,		
		criacao datetime DEFAULT CURRENT_TIMESTAMP,	
		alteracao datetime ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (id),
		FOREIGN KEY (integracao_id) REFERENCES integracoes(id)
	)`

	database.ExecutarQuery(db, sql)
}

// Gravar -
func (l Log) Gravar() bool {
	db := database.ObterConexao()
	defer db.Close()

	var sql string = `insert into logs 
		(api_destino, parametro, metodo, resposta, integracao_id) 
		values (?, ?, ?, ?, ?)`

	stmt, _ := db.Prepare(sql)

	resultado, err := stmt.Exec(
		l.APIDestino,
		l.Parametro,
		l.Metodo,
		l.Resposta,
		l.IntegracaoID)

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
func (l Log) Excluir() bool {
	db := database.ObterConexao()
	defer db.Close()

	stmt, _ := db.Prepare("DELETE FROM logs WHERE id = ?")
	var _, err = stmt.Exec(l.ID)

	if err != nil {
		return false
	}

	return true

}

// BuscarPorIDIntegracao -
func (l Log) BuscarPorIDIntegracao(idIntegracao int64) []Log {

	db := database.ObterConexao()
	defer db.Close()

	var sql string = `SELECT id, api_destino, metodo, parametro, resposta, integracao_id
		FROM logs WHERE integracao_id = ? ORDER BY id DESC`

	rows, _ := db.Query(sql, idIntegracao)
	defer rows.Close()

	var listaLog []Log
	for rows.Next() {

		var log Log

		rows.Scan(&log.ID,
			&log.APIDestino,
			&log.Metodo,
			&log.Parametro,
			&log.Resposta,
			&log.IntegracaoID)

		listaLog = append(listaLog, log)
	}

	return listaLog
}
