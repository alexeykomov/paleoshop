package controllers

import (
	"fmt"
	"net/http"
	"alexeykomov.me/paleoshop/storage"
)

func IndexHandler(response http.ResponseWriter, request *http.Request) {
	category, err := storage.GetRootCategory()
	if err != nil {
		response.WriteHeader(500)
		return
	}
	products, err := storage.GetProductsOfCategory([]string{}, category.Id)

	fmt.Fprintf(response, "<html><title>The first shop page</title></html>")
}
 
