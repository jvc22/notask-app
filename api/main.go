package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db/tasks.db")

	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS todos (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"title" TEXT,
		"completed" BOOLEAN
	);`
	
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Erro ao criar tabela:", err)
	}
	
	fmt.Println("Banco de dados inicializado com sucesso!")

	defer db.Close()
}