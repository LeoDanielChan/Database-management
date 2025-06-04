package generators

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"management.com/db"
	"management.com/utils"
)

const numClients = 100_000

func Client() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	dbDatos, err := db.Conectar("datos")
	if err != nil {
		log.Fatal("Conexión a base de datos 'datos' fallida: ", err)
	}
	defer dbDatos.Close()

	dbAirbus, err := db.Conectar("airbus380")
	if err != nil {
		log.Fatal("Conexión a base de datos 'airbus380' fallida: ", err)
	}
	defer dbAirbus.Close()

	nombres := utils.GetMames(dbDatos)
	apellidos := utils.GetLastName(dbDatos)
	estados := utils.GetStates(dbDatos)
	municipios := utils.GetMunicipalities(dbDatos)

	insert, err := dbAirbus.Prepare(`
		INSERT INTO clientes (cve_municipios, cve_estados, nombre, paterno, materno, fecha_nacimiento)
		VALUES (@p1, @p2, @p3, @p4, @p5, @p6)
	`)
	if err != nil {
		log.Fatal("Error al prepar insert:", err)
	}
	defer insert.Close()

	fmt.Println("Insertando clientes...")
	for i := 0; i < numClients; i++ {
		nombre := nombres[r.Intn(len(nombres))].Nombre
		paterno := apellidos[r.Intn(len(apellidos))].Apellido
		materno := apellidos[r.Intn(len(apellidos))].Apellido

		edad := r.Intn(86) + 5
		fechaNacimiento := time.Now().AddDate(-edad, -r.Intn(12), -r.Intn(28))

		estado := estados[r.Intn(len(estados))]
		var municipioID int
		for {
			m := municipios[r.Intn(len(municipios))]
			if m.CveEstado == estado.CveEstado {
				municipioID = m.CveMunicipio
				break
			}
		}

		_, err := insert.Exec(
			municipioID,
			estado.CveEstado,
			nombre,
			paterno,
			materno,
			fechaNacimiento,
		)
		if err != nil {
			log.Printf("Error al insertar cliente %d: %v", i+1, err)
		}

		if (i+1)%10000 == 0 {
			fmt.Printf("%d registros insertados...\n", i+1)
		}
	}
	fmt.Println("Proceso completado.")
}
