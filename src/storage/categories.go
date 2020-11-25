package storage

import (
	"fmt"
	"strconv"
	"strings"
)

type Category struct {
	id              int
	title           string
	coverURL        string
	childCategories []Category
	childProducts   []Product
	ancestorIds []int
}

type Product struct {
	id          int
	title       string
	coverURL    string
	description string
	pictures []Picture
}

type Picture struct {
	id int
	url string
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

func GetSubcategoriesOfCategory(categoryId string) ([]Category, error) {
	db, err := Connection()
	if err != nil {
		return []Category{}, err
	}
	rows, err := db.Query(
		fmt.Sprintf("select id, ancestry, title, cover from categories" +
			" where id %s", categoryId))
	var category Category
	var categories []Category
	var ancestry string
	for rows.Next() {
		err = rows.Scan(&category.id, &ancestry, &category.title, &category.coverURL)
		for _, id := range strings.Split(ancestry, "/") {
			idInt, err := strconv.Atoi(id)
			if err != nil {
				continue
			}
			category.ancestorIds = append(category.ancestorIds, idInt)
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func GetProductsOfCategory(ancestorIds []string, categoryId string) ([]Product, error) {
	db, err := Connection()
	if err != nil {
		return []Product{}, err
	}
	rows, err := db.Query(
		fmt.Sprintf("select products.id, products.title, products.cover "+
			"from categories, products, category_products cp where "+
			"categories.id = cp.category_id and "+
			"products.id = cp.product_id and ancestry like '/%s%%'",
			strings.Join(append(ancestorIds, categoryId), "/")))
	var product Product
	var products []Product
	for rows.Next() {
		err = rows.Scan(&product.id, &product.title, &product.coverURL)
		products = append(products, product)
	}
	return products, nil
}
