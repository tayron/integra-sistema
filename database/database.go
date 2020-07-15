package database

import (
	"database/sql"
	"fmt"

	// Uso indireto do drive mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/tayron/integra-sistema/configuration"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}

	return result
}

// ObterConexao - retorna conexão com banco de dados
func ObterConexao() *sql.DB {

	stringConexao := fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		configuration.GetConfiguracao("usuario"),
		configuration.GetConfiguracao("senha"),
		configuration.GetConfiguracao("localhost"),
		configuration.GetConfiguracao("porta"),
	)

	db, err := sql.Open("mysql", stringConexao)

	if err != nil {
		panic(err)
	}

	exec(db, "create database if not exists "+configuration.GetConfiguracao("banco"))
	exec(db, "use "+configuration.GetConfiguracao("banco"))

	return db
}
