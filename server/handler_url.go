package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) handlerURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world!")
}
