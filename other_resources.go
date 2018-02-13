package main

import (
	"encoding/json"
	"net/http"
)

type Version struct {
	Message  string `json:"message"`
	Subtitle string `json:"subtitle"`
}

type Operation struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Fields []Field `json:"fields"`
}

type Field struct {
    Name  string `json:"name"`
    Field string `json:"field"`
}

func GetVersion(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(Version{"Version 0.0.1", "FOAAS"})
}

func GetOperations(w http.ResponseWriter, r *http.Request) {
	var operations []Operation

	operations = append(operations, Operation{Name: "version", Url: "/version", Fields: []Field{}})
	operations = append(operations, Operation{Name: "Who the fuck are you anyway", Url: "/anyway/:company/:from", Fields: []Field{{Name: "Company", Field: "company"},{Name: "From", Field: "from"}}})
    json.NewEncoder(w).Encode(operations)
}