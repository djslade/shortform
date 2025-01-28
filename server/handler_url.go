package main

import (
	"fmt"
	"net/http"
)

func handlerURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world!")
}
