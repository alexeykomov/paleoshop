package main

import (
	"fmt"
	"net/http"
	"alexeykomov.me/paleoshop/controllers"
	"alexeykomov.me/paleoshop/storage"
)

func main() {
	http.HandleFunc("/", controllers.IndexHandler)
	db, err := storage.Connection()
	if err != nil {
		fmt.Println(err.Error())
	}
	storage.GetCategories()
	http.ListenAndServe(":8080", nil)
}
