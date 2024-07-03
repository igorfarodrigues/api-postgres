package models

import (
	"github.com/igorfarodrigues/api-postgres/db"
	"log"
)

func GetAll() (todos []Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Erro ao abrir conexão: %v", err)
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM todos`)
	if err != nil {
		log.Printf("Erro ao executar query: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			log.Printf("Erro ao fazer scan dos registros: %v", err)
			continue
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erro durante iteração dos registros: %v", err)
		return
	}

	return
}
