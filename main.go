package main

import (
	"log"
	"net/http"
)

func main() {
	router := GetRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
