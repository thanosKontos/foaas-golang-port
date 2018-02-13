package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type ApiResponse struct {
	Message  string `json:"message"`
	Subtitle string `json:"subtitle"`
}

func GetAnyway(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(ApiResponse{
		fmt.Sprintf("Who the fuck are you anyway, %s, why are you stirring up so much trouble, and, who pays you?", params["company"]),
		fmt.Sprintf("- %s", params["from"]),
	})
}
