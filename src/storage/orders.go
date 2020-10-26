package storage

func GetCategories() {
	db, err := Connection()
	if err == nil {
		db.Exec("SELECT * FROM CATEGORIES WHERE id = %id")	}
}
