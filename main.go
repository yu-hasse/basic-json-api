package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/user", makeHTTPHandleFunc(handleGetUserByID))
	http.ListenAndServe(":3000", nil)
}
