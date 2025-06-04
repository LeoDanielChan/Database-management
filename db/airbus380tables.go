package db

import (
	"fmt"
	"log"
)

func TablesAirbus380() {
	db, err := Conectar("airbus380")
	if err != nil {
		log.Fatal("Error abriendo la conexión:", err.Error())
	}
	defer db.Close()

	createTables := []string{
		`CREATE TABLE estados (
			cve_estados INT PRIMARY KEY,
			nombre VARCHAR(50),
			abreviatura VARCHAR(50)
		);`,

		`CREATE TABLE municipios (
			cve_municipios INT,
			cve_estados INT,
			nombre VARCHAR(50),
			PRIMARY KEY (cve_municipios, cve_estados),
			FOREIGN KEY (cve_estados) REFERENCES estados(cve_estados)
		);`,

		`CREATE TABLE clientes (
			cve_clientes INT IDENTITY(1,1) PRIMARY KEY,
			cve_municipios INT,
			cve_estados INT,
			nombre VARCHAR(50),
			paterno VARCHAR(50),
			materno VARCHAR(50),
			fecha_nacimiento DATETIME,
			FOREIGN KEY (cve_municipios, cve_estados) REFERENCES municipios(cve_municipios, cve_estados),
		);`,

		`CREATE TABLE detalle_vuelos (
			cve_detalle_vuelos INT IDENTITY(1,1) PRIMARY KEY,
			cve_vuelos INT,
			fecha_hora_salida DATETIME,
			capacidad INT,
			FOREIGN KEY (cve_vuelos) REFERENCES vuelos(cve_vuelos)
		);`,

		`CREATE TABLE ocupaciones (
			cve_ocupaciones INT IDENTITY(1,1) PRIMARY KEY,
			cve_detalle_vuelos INT,
			cve_clientes INT,
			FOREIGN KEY (cve_detalle_vuelos) REFERENCES detalle_vuelos(cve_detalle_vuelos),
			FOREIGN KEY (cve_clientes) REFERENCES clientes(cve_clientes)
		);`,
	}

	for _, table := range createTables {
		_, err := db.Exec(table)
		if err != nil {
			log.Fatalf("Error al crear tabla: \n %s\nDetalle: %v\n", table, err)
		} else {
			fmt.Printf("Tabla creada correctamente: \n %s \n", table)
		}
	}
}
