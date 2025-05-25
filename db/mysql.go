package db

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func Conectar(usuario, contrasena, host, puerto, dbname string) (*sql.DB, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", usuario, contrasena, host, puerto, dbname)
	return sql.Open("sqlserver", dsn)
}
