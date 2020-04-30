package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Database Connection String
func InitDB() (*sql.DB, error) {
	dbhost := os.Getenv("DBHOST")
	dbuser := os.Getenv("DBUSER")
	dbpass := os.Getenv("DBPASS")
	dbport := os.Getenv("DBPORT")
	dbname := os.Getenv("DB")
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp("+dbhost+":"+dbport+")/"+dbname)
	if err != nil {
		return nil, err
	}
	// db.SetMaxOpenConns(1000)
	// db.SetConnMaxLifetime(1000)
	// db.SetMaxIdleConns(10)
	return db, nil
}
