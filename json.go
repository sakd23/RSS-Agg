package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponsewithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Printf("Server error: %s", message)
	}
	//this Struct tells JSON msrshal that in json there will be an "error" field
	type ErrorResponse struct {
		Error string `json:"error"`
	}
	ResponsewithJSON(w, code, ErrorResponse{Error: message})
}

func ResponsewithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
