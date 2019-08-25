package controllers

import (
	"fmt"
	"net/http"

	"../utils"
)

//ProtectedEndpoint - Method for protectedEndpoint -> would accessed by jwt token
func (c Controller) ProtectedEndpoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// this would only be invoked only if the correct jwt token is passed
		fmt.Println("protectedEndpoint Invoked")
		utils.ResponseJSON(w, "Successfully invoked the protected endpoint.")
	}
}
