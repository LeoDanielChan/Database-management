package main

import (
	"fmt"
	"log"

	"management.com/db"
)

func main() {
	conexion, err := db.Conectar("leo10", "2004", "127.0.0.1", "1433", "datos")
	if err != nil {
		log.Fatal("Error conectando a MySQL:", err)
	}

	defer conexion.Close()

	// Probar conexión con Ping
	if err := conexion.Ping(); err != nil {
		log.Fatalf("❌ Falló el ping a la base de datos: %v", err)
	}
	fmt.Println("✅ Conexión exitosa a la base de datos.")

	query := "SELECT * FROM estados"
	filas, err := conexion.Query(query)
	if err != nil {
		log.Fatalf("❌ Error al ejecutar la consulta: %v", err)
	}
	defer filas.Close()

	columnas, err := filas.Columns()
	if err != nil {
		log.Fatalf("❌ Error obteniendo columnas: %v", err)
	}

	valores := make([]interface{}, len(columnas))
	punteros := make([]interface{}, len(columnas))
	for i := range valores {
		punteros[i] = &valores[i]
	}

	var resultados []map[string]interface{}

	fmt.Println("Resultado de la consulta:")
	for filas.Next() {
		err := filas.Scan(punteros...)
		if err != nil {
			log.Fatalf("❌ Error al leer fila: %v", err)
		}

		fila := make(map[string]interface{})

		for i, col := range columnas {
			fmt.Printf("%s: %v\n", col, valores[i])
		}

		resultados = append(resultados, fila)
	}

	fmt.Println("aqaa")
	fmt.Println(len(resultados))

	// fmt.Println("=== Crear Tabla Interactivamente ===")
	// if err := db.CrearTablaInteractiva(conexion); err != nil {
	// 	log.Fatal(err)
	// }
}
