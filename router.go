package main

import (
	"github.com/gorilla/mux"
	"github.com/thanosKontos/foaas-golang-port/api"
	"net/http"
)

func getRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/version", api.VersionHandler).Methods("GET")
	router.HandleFunc("/operations", api.OperationsHandler).Methods("GET")

	router.HandleFunc("/awesome/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/asshole/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/bag/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/because/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/bucket/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/bye/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/cool/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/cup/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/diabetes/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/everyone/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/everything/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/family/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/fascinating/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/flying/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/fyyff/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/give/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/horse/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/immensity/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/life/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/looking/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/maybe/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/me/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/mornin/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/no/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/pink/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/programmer/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/retard/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/ridiculous/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/rtfm/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/sake/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/shit/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/single/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/thanks/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/that/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/this/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/too/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/tucker/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/what/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/zayn/{first}", api.OneParameterHandler).Methods("GET")
	router.HandleFunc("/zero/{first}", api.OneParameterHandler).Methods("GET")

	router.HandleFunc("/anyway/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/anyway/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/anyway/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/anyway/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/back/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/bday/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/blackadder/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/bus/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/caniuse/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/chainsaw/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/cocksplat/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/dalton/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/deraadt/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/donut/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/gfy/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/ing/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/keep/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/keepcalm/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/king/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/linus/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/look/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/nugget/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/off/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/off-with/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/outside/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/particular/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/problem/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/pulp/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/shakespeare/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/shutup/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/thinking/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/thumbs/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/xmas/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/yoda/{first}/{second}", api.TwoParametersHandler).Methods("GET")
	router.HandleFunc("/you/{first}/{second}", api.TwoParametersHandler).Methods("GET")

	router.HandleFunc("/dosomething/{first}/{second}/{third}", api.ThreeParametersHandler).Methods("GET")
	router.HandleFunc("/ballmer/{first}/{second}/{third}", api.ThreeParametersHandler).Methods("GET")

	router.HandleFunc("/field/{first}/{second}/{third}", api.IrregularRequestHandler).Methods("GET")
	router.HandleFunc("/greed/{first}/{second}", api.IrregularRequestHandler).Methods("GET")
	router.HandleFunc("/madison/{first}/{second}", api.IrregularRequestHandler).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(api.GenericErrorHandler)
	router.MethodNotAllowedHandler = http.HandlerFunc(api.GenericErrorHandler)

	return router
}
