package main

import (
	"log"
	"net/http"
	"fmt"
	"test_bar/handlers"
)

func main() {
	http.Handle("/bar", handlers.NewBarHandler())

	port := 8081
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}