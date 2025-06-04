package generators

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"management.com/db"
)

func Flight() {
	db, err := db.Conectar("airbus380")
	if err != nil {
		log.Fatal("Conexión a base de datos 'airbus380' fallida: ", err)
	}
	defer db.Close()

	// 1. Obtener aeropuertos de México
	aeropuertosMex, err := obtenerAeropuertosMexicanos(db)
	if err != nil {
		log.Fatal(err)
	}

	// 2. Obtener vuelos válidos
	vuelos, err := obtenerVuelosConAeropuertos(db, aeropuertosMex)
	if err != nil {
		log.Fatal(err)
	}

	if len(vuelos) == 0 {
		log.Fatal("No hay vuelos válidos desde o hacia México")
	}

	// 3. Insertar registros en detalle_vuelos
	err = insertarDetalleVuelos(db, vuelos, 2000)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ 2000 registros insertados en detalle_vuelos")
}

func obtenerAeropuertosMexicanos(db *sql.DB) ([]int, error) {
	query := `
		SELECT a.cve_aeropuertos
		FROM aeropuertos a
		JOIN ciudades c ON a.cve_ciudades = c.cve_ciudades
		JOIN paises p ON c.cve_paises = p.cve_paises
		WHERE p.nombre = 'México'
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener aeropuertos de México: %w", err)
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func obtenerVuelosConAeropuertos(db *sql.DB, aeropuertos []int) ([]int, error) {
	// Construir cláusula IN
	inClause := "("
	for i, id := range aeropuertos {
		if i > 0 {
			inClause += ","
		}
		inClause += fmt.Sprintf("%d", id)
	}
	inClause += ")"

	query := fmt.Sprintf(`
		SELECT cve_vuelos
		FROM vuelos
		WHERE cve_aeropuertos__origen IN %s OR cve_aeropuertos__destino IN %s
	`, inClause, inClause)

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener vuelos válidos: %w", err)
	}
	defer rows.Close()

	var vuelos []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		vuelos = append(vuelos, id)
	}
	return vuelos, nil
}

func insertarDetalleVuelos(db *sql.DB, vuelos []int, cantidad int) error {
	stmt, err := db.Prepare(`
		INSERT INTO detalle_vuelos (cve_vuelos, fecha_hora_salida, capacidad)
		VALUES (@p1, @p2, @p3)
	`)
	if err != nil {
		return fmt.Errorf("error preparando insert: %w", err)
	}
	defer stmt.Close()

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < cantidad; i++ {
		vueloID := vuelos[rand.Intn(len(vuelos))]
		fecha := fechaAleatoria2023()

		_, err := stmt.Exec(vueloID, fecha, 500)
		if err != nil {
			log.Printf("Error insertando registro %d: %v\n", i, err)
		}
	}
	return nil
}

func fechaAleatoria2023() time.Time {
	// Año 2023
	fecha := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	dias := rand.Intn(365)
	horas := rand.Intn(24) // Hora exacta
	return fecha.AddDate(0, 0, dias).Add(time.Duration(horas) * time.Hour)
}

func capacidadAleatoria() int {
	return (rand.Intn(4) + 7) * 50 // 350, 400, 450, 500
}
