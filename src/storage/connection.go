package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var globalDB *sql.DB = nil

const (
	host     = "localhost"
	port     = 5432
	user     = "paleoshop_user"
	password = "JoToo8kZzhHQxs6H"
	dbname   = "postgres"
)

func params() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s " +
		"dbname=%s sslmode=disable", host, port, user, password, dbname)
}

func Connection() (*sql.DB, error) {
	if globalDB != nil {
		return globalDB, nil
	}

	db, err := sql.Open("postgres", params())
	if err == nil {
		err = db.Ping()
		if err != nil {
			return nil, err
		}
		//defer db.Close()
		globalDB = db
		return db, nil
	}
	return nil, err
}
