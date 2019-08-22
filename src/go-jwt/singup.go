package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"golang.org/x/crypto/bcrypt"
)

//Method to handle the singup request
func singup(w http.ResponseWriter, r *http.Request) {

	var user User
	var error Error
	//New Decoder returns a decoder that reads from r -> request
	json.NewDecoder(r.Body).Decode(&user)
	//spew helps in logging detailed values of the struct
	spew.Dump(user)

	//Vaidations
	if user.Email == "" {
		error.Message = "Email is missing"
		respondError(w, http.StatusBadRequest, error)
		return
	}
	if user.Password == "" {
		error.Message = "Password is missing"
		respondError(w, http.StatusBadRequest, error)
		return
	}

	//As the validations are successful
	// Hash/encrypt the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal("Error by encrypting the password")
	}

	user.Password = string(hash) //assigning the hash to password after converting back to string
	//inserting the record into DB
	query := "insert into users (email, password) values ($1, $2) RETURNING id;"
	//Scan is used as we expect the query to return ID -> scan returns err or, nil
	err = db.QueryRow(query, user.Email, user.Password).Scan(&user.ID)

	if err != nil {
		error.Message = "Internal Server Error"
		respondError(w, http.StatusInternalServerError, error)
		return
	}

	//For success - if no error the record is inserted successfully
	user.Password = ""                                 // to not expose the password
	w.Header().Set("Content-Type", "application/json") //set hearer
	responseJSON(w, user)                              //util method
}
