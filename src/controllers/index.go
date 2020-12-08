package controllers

import (
	"fmt"
	"io"
	"net/http"
	"alexeykomov.me/paleoshop/storage"
	"html/template"
	"os"
)

type IndexPageData struct {
	Categories []storage.Category
	Products []storage.Product
}

func IndexHandler(response http.ResponseWriter, request *http.Request) {
	rootCategory, err := storage.GetRootCategory()
	fmt.Printf("rootCategory: %v", rootCategory)
	if err != nil {
		response.WriteHeader(500)
		io.WriteString(response, err.Error())
		return
	}
	categories, err := storage.GetSubcategoriesOfCategory(rootCategory.Id)
	products, err := storage.GetProductsOfCategory(rootCategory)
	wd, err := os.Getwd()
	if err != nil {
		response.WriteHeader(500)
		io.WriteString(response, err.Error())
	}
	indexPage := template.Must(template.ParseFiles(wd + "/views/index.gohtml"))
	fmt.Printf("%v", categories)
	indexPage.Execute(response, IndexPageData{
		Categories: categories,
		Products: products,
	})
}

