package main

import (
	"net/http"
	"alexeykomov.me/shop2000/controllers"
)

func main() {
	http.HandleFunc("/", controllers.IndexHandler)
	http.ListenAndServe(":8080", nil)
}
