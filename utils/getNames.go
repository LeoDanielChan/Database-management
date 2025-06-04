package utils

import (
	"database/sql"
	"log"
)

type Nombre struct {
	Nombre string
}

func GetMames(db *sql.DB) []Nombre {
	rows, err := db.Query("SELECT nombre FROM nombres")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var lista []Nombre
	for rows.Next() {
		var n Nombre
		if err := rows.Scan(&n.Nombre); err == nil {
			lista = append(lista, n)
		}
	}
	return lista
}
