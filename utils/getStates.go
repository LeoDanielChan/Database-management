package utils

import (
	"database/sql"
	"log"
)

type Estado struct {
	CveEstado int
}

func GetStates(db *sql.DB) []Estado {
	rows, err := db.Query("SELECT cve_estado FROM estados")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var lista []Estado
	for rows.Next() {
		var e Estado
		if err := rows.Scan(&e.CveEstado); err == nil {
			lista = append(lista, e)
		}
	}
	return lista
}
