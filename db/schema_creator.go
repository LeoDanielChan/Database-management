package db

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func CrearTablaInteractiva(db *sql.DB) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nombre de la tabla: ")
	nombreTabla, _ := reader.ReadString('\n')
	nombreTabla = strings.TrimSpace(nombreTabla)

	campos := []string{}
	for {
		fmt.Print("Nombre del campo (deja vacío para terminar): ")
		nombreCampo, _ := reader.ReadString('\n')
		nombreCampo = strings.TrimSpace(nombreCampo)

		if nombreCampo == "" {
			break
		}

		fmt.Print("Tipo de dato (ej. VARCHAR(100), INT, DATE): ")
		tipoDato, _ := reader.ReadString('\n')
		tipoDato = strings.TrimSpace(tipoDato)

		campos = append(campos, fmt.Sprintf("%s %s", nombreCampo, tipoDato))
	}

	if len(campos) == 0 {
		fmt.Println("No se definieron campos. Cancelando creación.")
		return nil
	}

	query := fmt.Sprintf("CREATE TABLE %s (%s)", nombreTabla, strings.Join(campos, ", "))
	fmt.Println("\nConsulta generada:")
	fmt.Println(query)

	fmt.Print("¿Deseas ejecutar esta consulta? (s/n): ")
	confirmacion, _ := reader.ReadString('\n')
	confirmacion = strings.TrimSpace(confirmacion)

	if strings.ToLower(confirmacion) == "s" {
		_, err := db.Exec(query)
		if err != nil {
			return fmt.Errorf("error al crear la tabla: %w", err)
		}
		fmt.Println("✅ Tabla creada con éxito.")
	} else {
		fmt.Println("❌ Creación cancelada.")
	}

	return nil
}
