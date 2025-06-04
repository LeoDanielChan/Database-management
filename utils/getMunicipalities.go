package utils

import (
	"database/sql"
	"log"
)

type Municipio struct {
	CveMunicipio int
	CveEstado    int
}

func GetMunicipalities(db *sql.DB) []Municipio {
	rows, err := db.Query("SELECT cve_municipios, cve_estados FROM municipios")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var lista []Municipio
	for rows.Next() {
		var m Municipio
		if err := rows.Scan(&m.CveMunicipio, &m.CveEstado); err == nil {
			lista = append(lista, m)
		}
	}
	return lista
}
