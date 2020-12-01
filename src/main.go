package main

import (
	"alexeykomov.me/paleoshop/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.IndexHandler)
	http.ListenAndServe(":8080", nil)
}
