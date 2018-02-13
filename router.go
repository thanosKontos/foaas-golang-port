package main

import (
    "github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/version", GetVersion).Methods("GET")
	router.HandleFunc("/operations", GetOperations).Methods("GET")
	router.HandleFunc("/anyway/{company}/{from}", GetAnyway).Methods("GET")

	return router
}