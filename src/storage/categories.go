package storage

import "database/sql"

func GetCategories() (sql.Result, error) {
	db, err := Connection()
	if err != nil {
		return nil, err
	}
	rows, err := db.Exec("SELECT * FROM CATEGORIES WHERE ancestry = ''")
	rows
}
