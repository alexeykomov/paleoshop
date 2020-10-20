package controllers

import (
	"fmt"
	"net/http"
)

func IndexHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "<html><title>The first shop page</title></html>")
}
