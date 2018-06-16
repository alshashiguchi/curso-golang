package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:1234567@/cursogo")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Inicia a transação
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios (id, nome) values (?, ?)")

	stmt.Exec(2000, "Arthur")
	stmt.Exec(2001, "Yoshiharu")
	_, err = stmt.Exec(1, "Andre") // chave duplicada

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
