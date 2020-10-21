package storage

import (
	"database/sql"
)

func Connection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "")
	if err == nil {
		err = db.Ping()
		if err != nil {
			return nil, err
		}
		defer db.Close()
		return db, nil
	}
	return nil, err
}
