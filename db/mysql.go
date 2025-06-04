package db

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func Conectar(dbname string) (*sql.DB, error) {
	dsn := fmt.Sprintf("sqlserver://leo10:2004@127.0.0.1:1433?database=%s", dbname)
	return sql.Open("sqlserver", dsn)
}
