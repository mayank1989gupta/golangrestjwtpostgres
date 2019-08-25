package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"../models"
	"github.com/dgrijalva/jwt-go"
)

//RespondError - function to create response for error
func RespondError(w http.ResponseWriter, status int, error models.Error) {
	w.WriteHeader(status)                              //HTTP Status
	w.Header().Set("Content-Type", "application/json") //set header
	json.NewEncoder(w).Encode(error)
}

//ResponseJSON - func to write the response JSON
func ResponseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json") //set header
	json.NewEncoder(w).Encode(data)
}

//GenerateToken - Method to generate token
func GenerateToken(user models.User) (string, error) {
	secret := os.Getenv("APP_DB_SECRET") //secret of the token
	//first param is the algorithm, second: struct of the claims we want to use
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email, //Email
		"iss":   "course",   //Issuer
	})
	//To generate toekn string for the clients
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}
	return tokenString, err
}
