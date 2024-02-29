package main

import (
	"net/http"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

func main()  {
	http.HandleFunc("/user", makeHttpHandler(handleGetUserById))
	http.ListenAndServe(":3000", nil)
}