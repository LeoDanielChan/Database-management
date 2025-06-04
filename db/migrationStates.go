package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

type Estado struct {
	CveEstado   int
	Nombre      string
	Abreviatura string
}

type Municipio struct {
	CveMunicipio int
	CveEstado    int
	Nombre       string
}

func MigrationStates() {
	dbAirbus, err := Conectar("airbus380")
	if err != nil {
		log.Fatal("Error abriendo la conexión:", err.Error())
	}
	defer dbAirbus.Close()

	dbDatos, err := Conectar("datos")
	if err != nil {
		log.Fatal("Conexión a base de datos 'datos' fallida: ", err)
	}
	defer dbDatos.Close()

	migrarEstados(dbDatos, dbAirbus)
	migrarMunicipios(dbDatos, dbAirbus)

	fmt.Println("Migración completada correctamente.")
}

func migrarEstados(origen, destino *sql.DB) {
	rows, err := origen.Query("SELECT * FROM estados")
	if err != nil {
		log.Fatal("Error leyendo estados:", err)
	}
	defer rows.Close()

	stmt, err := destino.Prepare(`
		INSERT INTO estados (cve_estados, nombre, abreviatura)
		VALUES (@p1, @p2, @p3)
	`)
	if err != nil {
		log.Fatal("Error preparando insert estados:", err)
	}
	defer stmt.Close()

	for rows.Next() {
		var e Estado
		if err := rows.Scan(&e.CveEstado, &e.Nombre, &e.Abreviatura); err != nil {
			log.Println("Error escaneando estado:", err)
			continue
		}

		_, err := stmt.Exec(e.CveEstado, e.Nombre, e.Abreviatura)
		if err != nil {
			log.Println("Error insertando estado:", err)
		}
	}
	fmt.Println("Estados migrados correctamente.")
}

func migrarMunicipios(origen, destino *sql.DB) {
	rows, err := origen.Query("SELECT * FROM municipios")
	if err != nil {
		log.Fatal("Error leyendo municipios:", err)
	}
	defer rows.Close()

	stmt, err := destino.Prepare(`
		INSERT INTO municipios (cve_municipios, cve_estados, nombre)
		VALUES (@p1, @p2, @p3)
	`)
	if err != nil {
		log.Fatal("Error preparando insert municipios:", err)
	}
	defer stmt.Close()

	for rows.Next() {
		var m Municipio
		if err := rows.Scan(&m.CveMunicipio, &m.CveEstado, &m.Nombre); err != nil {
			log.Println("Error escaneando municipio:", err)
			continue
		}

		_, err := stmt.Exec(m.CveMunicipio, m.CveEstado, m.Nombre)
		if err != nil {
			log.Println("Error insertando municipio:", err)
		}
	}
	fmt.Println("Municipios migrados correctamente.")
}
