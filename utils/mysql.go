package utils

import (
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver

	"github.com/jmoiron/sqlx"
)

// InitMySQL Function that initializes mysql
func InitMySQL(dbName string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", "root:root@tcp(127.0.0.1:3306)/"+dbName)
	if err != nil {
		log.Fatalln(err)
	}

	return db, err
}
