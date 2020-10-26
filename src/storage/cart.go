package storage

func GetCart() {
	db, err := Connection()
	if err == nil {
		db.Exec("SELECT * FROM CATEGORIES WHERE id = %id")	}
}
