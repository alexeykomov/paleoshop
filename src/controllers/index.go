package controllers

import (
	"io"
	"net/http"
	"alexeykomov.me/paleoshop/storage"
	"html/template"
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

	indexPage := template.Must(template.ParseFiles("views/index.html"))
	indexPage.Execute(response, IndexPageData{
		Categories: categories,
		Products: products,
	})
}

