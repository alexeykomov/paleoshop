package storage

import (
	"fmt"
	"strconv"
	"strings"
)

type Category struct {
	Id              int
	Title           string
	coverURL        string
	childCategories []Category
	childProducts   []Product
	ancestorIds     []int
}

type Product struct {
	id          int
	title       string
	coverURL    string
	description string
	pictures    []Picture
}

type Picture struct {
	id  int
	url string
}

func GetRootCategory() (Category, error) {
	db, err := Connection()
	if err != nil {
		return Category{}, err
	}
	row := db.QueryRow("select id, ancestry, title, cover from " +
		"categories where ancestry = ''")
	var rootCategory Category
	var ancestry string
	err = row.Scan(&rootCategory.Id, &ancestry,
		&rootCategory.Title, &rootCategory.coverURL)
	populateAncestry(ancestry, rootCategory)
	if err != nil {
		return Category{}, err
	}
	return rootCategory, nil
}

func GetSubcategoriesOfCategory(categoryId int) ([]Category, error) {
	fmt.Printf("categoryId: %d", categoryId)
	db, err := Connection()
	if err != nil {
		return []Category{}, err
	}
	rows, err := db.Query(
		fmt.Sprintf("select id, ancestry, title, cover from "+
			"categories where ancestry = '%s'", strconv.Itoa(categoryId)))
	if err != nil {
		return []Category{}, err
	}
	var category Category
	var categories []Category
	var ancestry string
	for rows.Next() {
		err = rows.Scan(&category.Id, &ancestry, &category.Title, &category.coverURL)
		populateAncestry(ancestry, category)
		categories = append(categories, category)
	}
	return categories, nil
}

func GetProductsOfCategory(category Category) ([]Product, error) {
	db, err := Connection()
	if err != nil {
		return []Product{}, err
	}
	ancestry := getAncestry(category)
	fmt.Printf("ancestry: %s\n", ancestry)
	query := fmt.Sprintf(
		"select products.id, products.title, products.cover "+
			"from categories, products, category_products cp where "+
			"categories.id = cp.category_id and "+
			"products.id = cp.product_id and "+
			"(ancestry = '%s' or ancestry like '%s/%%')",
		strconv.Itoa(category.Id), ancestry)
	fmt.Printf("Query: %s\n", query)
	rows, err := db.Query(query)
	if err != nil {
		return []Product{}, err
	}
	var product Product
	var products []Product
	for rows.Next() {
		err = rows.Scan(&product.id, &product.title, &product.coverURL)
		products = append(products, product)
	}
	return products, nil
}

func getAncestry(category Category) string {
	var ancestorIdsStr []string
	for _, id := range category.ancestorIds {
		ancestorIdsStr = append(ancestorIdsStr, strconv.Itoa(id))
	}
	return strings.Join(ancestorIdsStr, "/")
}

func populateAncestry(ancestry string, category Category) {
	for _, id := range strings.Split(ancestry, "/") {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			continue
		}
		category.ancestorIds = append(category.ancestorIds, idInt)
	}
}
