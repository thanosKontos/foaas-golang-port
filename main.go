package main

import (
    "encoding/json"
    "log"
	"net/http"
)

func genericError(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ApiResponse{
		"622 - All The Fucks",
		"Server invites you to consider the truly monumental amount of fucks it couldn't give about your request.",
	})
}

func main() {
	router := GetRouter()
	router.NotFoundHandler = http.HandlerFunc(genericError)
	router.MethodNotAllowedHandler = http.HandlerFunc(genericError)

    log.Fatal(http.ListenAndServe(":8000", router))
}