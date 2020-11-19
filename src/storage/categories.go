package storage

import (
	"fmt"
	"strings"
)

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

func GetSubcategoriesOfCategory(ancestorAndCurrentIds []string) ([]Category, error) {
	db, err := Connection()
	if err != nil {
		return []Category{}, err
	}
	rows, err := db.Query(
		fmt.Sprintf("select products.id, ancestry, products.title "+
			"from categories, products, category_products cp where "+
			"categories.id = cp.category_id and "+
			"products.id = cp.product_id and ancestry like '/%s%%'", strings.Join(ancestorAndCurrentIds, "/")))
	var category Category
	var categories []Category
	for rows.Next() {
		err = rows.Scan(&category.id, &category.title, &category.coverURL)
		categories = append(categories, category)
	}
	return categories, nil
}
