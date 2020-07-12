package database

import (
	"database/sql"

	// Uso indireto do drive mysql
	_ "github.com/go-sql-driver/mysql"
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
	db, err := sql.Open("mysql", "root:yakTLS&70c52@tcp(172.30.0.2:3306)/")

	if err != nil {
		panic(err)
	}

	exec(db, "create database if not exists cursogo")
	exec(db, "use cursogo")

	return db
}
