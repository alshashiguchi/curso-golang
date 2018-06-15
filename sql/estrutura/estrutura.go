package main

import (
	"database/sql"
	// Não vamos utilizar explicitamente o pacote
	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "root:1234567@/") // quando passo o mysql ele busca o indiretamente o driver que é o go-sql-driver/mysql

	if err != nil {
		panic(err)
	}

	defer db.Close()

	exec(db, "create database if not exists cursogo")
	exec(db, "use cursogo")
	exec(db, "drop table if exists usuarios")
	exec(db, `create table usuarios (
		id integer auto_increment,
		nome varchar(80),

		PRIMARY KEY (id)
	)`)
}
