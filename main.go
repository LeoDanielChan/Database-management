package main

import (
	"management.com/db"
	"management.com/generators"
)

func main() {
	db.TablesAirbus380()
	db.MigrationStates()
	generators.Client()
	generators.Flight()
	generators.Occupations()
}
