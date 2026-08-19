package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

func CreateDbSqlLite(path string) (*sql.DB, error) {
	db, errDb := sql.Open("sqlite", path)
	if errDb != nil {
		return nil, errDb
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxIdleTime(1 * time.Minute)
	db.SetConnMaxLifetime(5 * time.Hour)

	if db.Ping() != nil {
		return nil, fmt.Errorf("Error al ping de la db")
	}

	return db, nil
}
