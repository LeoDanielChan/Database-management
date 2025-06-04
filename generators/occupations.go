package generators

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"management.com/db"
)

type Vuelo struct {
	CveDetalleVuelo int
	Capacidad       int
}

func Occupations() {
	db, err := db.Conectar("airbus380")
	if err != nil {
		log.Fatal("Conexión a base de datos 'airbus380' fallida: ", err)
	}
	defer db.Close()

	clientes, err := obtenerClientes(db)
	if err != nil {
		log.Fatal(err)
	}

	vuelos, err := obtenerDetalleVuelos(db)
	if err != nil {
		log.Fatal(err)
	}

	total := 0
	batchSize := 1000
	target := 1_000_000
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	tx, _ := db.Begin()
	if err != nil {
		log.Fatal("Error al iniciar transacción:", err)
	}

	stmt, _ := tx.Prepare("INSERT INTO ocupaciones (cve_clientes, cve_detalle_vuelos) VALUES (@p1, @p2)")
	if err != nil {
		log.Fatal("Error al preparar statement:", err)
	}
	defer stmt.Close()

	fmt.Println("Iniciando inserción de ocupaciones...")

	for total < target {
		// Seleccionar un vuelo aleatorio
		vuelo := vuelos[r.Intn(len(vuelos))]

		// Insertar hasta llenar la capacidad del vuelo (500)
		ocupaciones := min(500, target-total)
		for i := 0; i < ocupaciones; i++ {
			cliente := clientes[r.Intn(len(clientes))]
			_, err := stmt.Exec(cliente, vuelo.CveDetalleVuelo)
			if err != nil {
				log.Printf("Error insertando ocupación: %v\n", err)
				continue
			}
			total++

			// Commit parcial cada batchSize
			if total%batchSize == 0 {
				tx.Commit()
				tx, err = db.Begin()
				if err != nil {
					log.Fatal("Error al iniciar nueva transacción:", err)
				}
				stmt, err = tx.Prepare("INSERT INTO ocupaciones (cve_clientes, cve_detalle_vuelos) VALUES (@p1, @p2)")
				if err != nil {
					log.Fatal("Error al preparar statement:", err)
				}
			}

			// Mostrar progreso cada 10,000 registros
			if total%10000 == 0 {
				fmt.Printf("\rProgreso: %d/%d (%.1f%%)", total, target, float64(total)/float64(target)*100)
			}
		}
	}

	// Commit final
	tx.Commit()
	fmt.Printf("\n✅ Inserción completada: %d registros insertados en ocupaciones\n", total)
}

func obtenerClientes(db *sql.DB) ([]int, error) {
	rows, err := db.Query("SELECT cve_clientes FROM clientes")
	if err != nil {
		return nil, fmt.Errorf("error al obtener clientes: %w", err)
	}
	defer rows.Close()

	var clientes []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		clientes = append(clientes, id)
	}
	return clientes, nil
}

func obtenerDetalleVuelos(db *sql.DB) ([]Vuelo, error) {
	rows, err := db.Query("SELECT cve_detalle_vuelos, capacidad FROM detalle_vuelos")
	if err != nil {
		return nil, fmt.Errorf("error al obtener detalle_vuelos: %w", err)
	}
	defer rows.Close()

	var vuelos []Vuelo
	for rows.Next() {
		var v Vuelo
		if err := rows.Scan(&v.CveDetalleVuelo, &v.Capacidad); err != nil {
			return nil, err
		}
		vuelos = append(vuelos, v)
	}
	return vuelos, nil
}
