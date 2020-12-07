package controllers

import (
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
	if err != nil {
		response.WriteHeader(500)
		io.WriteString(response, err.Error())
		return
	}
	categories, err := storage.GetSubcategoriesOfCategory([]string{}, rootCategory.Id)
	products, err := storage.GetProductsOfCategory([]string{}, rootCategory.Id)
	wd, err := os.Getwd()
	if err != nil {
		response.WriteHeader(500)
		io.WriteString(response, err.Error())
	}
	indexPage := template.Must(template.ParseFiles(wd + "/views/index.gohtml"))
	indexPage.Execute(response, IndexPageData{
		Categories: categories,
		Products: products,
	})
}

