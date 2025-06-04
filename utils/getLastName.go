package utils

import (
	"database/sql"
	"log"
)

type Apellido struct {
	Apellido string
}

func GetLastName(db *sql.DB) []Apellido {
	rows, err := db.Query("SELECT apellido FROM apellidos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var lista []Apellido
	for rows.Next() {
		var a Apellido
		if err := rows.Scan(&a.Apellido); err == nil {
			lista = append(lista, a)
		}
	}
	return lista
}
