package main

import (
	"encoding/json"
	"net/http"
)

//function to create response for error
func respondError(w http.ResponseWriter, status int, error Error) {
	w.WriteHeader(status) //HTTP Status
	json.NewEncoder(w).Encode(error)
}

//func to write the response JSON
func responseJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}
