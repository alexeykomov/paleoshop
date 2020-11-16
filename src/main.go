package main

import (
	"fmt"
	"net/http"
	"alexeykomov.me/paleoshop/controllers"
	"alexeykomov.me/paleoshop/storage"
)

func main() {
	http.HandleFunc("/", controllers.IndexHandler)
	_, err := storage.Connection()
	if err != nil {
		fmt.Println(err.Error())
	}
	http.ListenAndServe(":8080", nil)
}
