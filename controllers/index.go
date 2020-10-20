package main

import (
	"fmt"
	"net/http"
)

func indexHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "<html><title>The first shop page</title></html>")
}
