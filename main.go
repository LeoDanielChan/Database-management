package main

import (
	"fmt"
	"log"

	"management.com/db"
)

func main() {
	conexion, err := db.Conectar("usuario", "contrasena", "127.0.0.1", "3306", "nombre_base_datos")
	if err != nil {
		fmt.Println("a")
		log.Fatal("Error conectando a MySQL:", err)
	}
	defer conexion.Close()

	fmt.Println("=== Crear Tabla Interactivamente ===")
	if err := db.CrearTablaInteractiva(conexion); err != nil {
		log.Fatal(err)
	}
}
