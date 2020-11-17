package storage

import "database/sql"

type Category struct {
	id              int
	title           string
	coverURL        string
	childCategories []Category
	childProducts   []product
}

type product struct {
	id          int
	title       string
	coverURL    string
	description string
}

func GetRootCategory() (Category, error) {
	db, err := Connection()
	if err != nil {
		return Category{}, err
	}
	row := db.QueryRow("SELECT id, title, cover FROM CATEGORIES " +
		"WHERE ancestry = ''")
	var rootCategory Category
	err = row.Scan(&rootCategory.id, &rootCategory.title, &rootCategory.coverURL)
	if err != nil {
		return Category{}, err
	}
	return rootCategory, nil
}

func GetSubcategoriesOfCategory(categoryId int) (Category, error) {
	db, err := Connection()
	if err != nil {
		return Category{}, err
	}
	rows, err := db.Query("SELECT * FROM CATEGORIES " +
		"WHERE ancestry LIKE %id%")
	var category Category
	for rows.Next() {
		err = rows.Scan(&category.id, &category.title, &category.coverURL)
	}
}
