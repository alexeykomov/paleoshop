package storage

struct User {
	name: String
	email: String
}

func GetUsers() {
	db, err := Connection()
	if err == nil {
		db.Exec("SELECT * FROM USER WHERE id = %id")	}
}
