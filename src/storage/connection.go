package storage

import (
	"database/sql"
)

var globalDB *sql.DB = nil

func Connection() (*sql.DB, error) {
	if globalDB != nil {
		return globalDB, nil
	}

	db, err := sql.Open("postgres", "")
	if err == nil {
		err = db.Ping()
		if err != nil {
			return nil, err
		}
		defer db.Close()
		globalDB = db
		return db, nil
	}
	return nil, err
}
