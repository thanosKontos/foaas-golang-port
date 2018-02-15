package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func genericError(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ApiResponse{
		"622 - All The Fucks",
		"Server invites you to consider the truly monumental amount of fucks it couldn't give about your request.",
	})
}

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/version", GetVersion).Methods("GET")
	router.HandleFunc("/operations", GetOperations).Methods("GET")
	router.HandleFunc("/anyway/{company}/{from}", GetCompanyFromHandler).Methods("GET")
	router.HandleFunc("/awesome/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/asshole/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/bag/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/because/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/bucket/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/bye/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/cool/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/cup/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/diabetes/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/everyone/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/everything/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/family/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/fascinating/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/flying/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/fyyff/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/give/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/horse/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/immensity/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/life/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/looking/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/maybe/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/me/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/mornin/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/no/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/pink/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/programmer/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/retard/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/ridiculous/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/rtfm/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/sake/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/shit/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/single/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/thanks/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/that/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/this/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/too/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/tucker/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/what/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/zayn/{from}", GetFromHandler).Methods("GET")
	router.HandleFunc("/zero/{from}", GetFromHandler).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(genericError)
	router.MethodNotAllowedHandler = http.HandlerFunc(genericError)

	return router
}
