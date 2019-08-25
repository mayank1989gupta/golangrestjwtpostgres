package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"../models"
	userrepository "../repository/user"
	"../utils"

	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

//Controller - To pass db which is used inside the login and signup methods
type Controller struct{}

// Singup - Sing up returns the http.HandlerFunc &, the func literal is the handler func returns the same
func (c Controller) Singup(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user models.User
		var error models.Error
		//New Decoder returns a decoder that reads from r -> request
		json.NewDecoder(r.Body).Decode(&user)
		//spew helps in logging detailed values of the struct
		spew.Dump(user)

		//Vaidations
		if user.Email == "" {
			error.Message = "Email is missing"
			utils.RespondError(w, http.StatusBadRequest, error)
			return
		}
		if user.Password == "" {
			error.Message = "Password is missing"
			utils.RespondError(w, http.StatusBadRequest, error)
			return
		}

		//As the validations are successful
		// Hash/encrypt the password
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			log.Fatal("Error by encrypting the password")
		}

		user.Password = string(hash) //assigning the hash to password after converting back to string
		userRepository := userrepository.UserRepository{}
		user = userRepository.Signup(db, user)
		if err != nil {
			error.Message = "Internal Server Error"
			utils.RespondError(w, http.StatusInternalServerError, error)
			return
		}

		//For success - if no error the record is inserted successfully
		user.Password = ""                                 // to not expose the password
		w.Header().Set("Content-Type", "application/json") //set hearer
		utils.ResponseJSON(w, user)                        //util method
	}
}

//Login - Method to handle the login request
func (c Controller) Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		var jwt models.JWT
		var error models.Error
		//decode the user body and retrieve user
		json.NewDecoder(r.Body).Decode(&user)
		//Validations
		if user.Email == "" {
			error.Message = "Email is missing"
			utils.RespondError(w, http.StatusBadRequest, error)
			return
		}
		if user.Password == "" {
			error.Message = "Password is missing"
			utils.RespondError(w, http.StatusBadRequest, error)
			return
		}

		//check if user is in DB
		password := user.Password
		userRepository := userrepository.UserRepository{}
		user, err := userRepository.Login(db, user) // call to repository

		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "User does not exist"
				utils.RespondError(w, http.StatusBadRequest, error)
				return
			}
			log.Fatal(err) //if not returned from if
		}
		//check password with the hashed stored password
		hashedPassword := user.Password
		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

		if err != nil {
			error.Message = "Password does not match"
			utils.RespondError(w, http.StatusBadRequest, error)
			return
		}

		//Generating token
		token, err := utils.GenerateToken(user)
		if err != nil {
			error.Message = "Error while generating token"
			utils.RespondError(w, http.StatusInternalServerError, error)
			return
		}

		//As everything is ok
		w.WriteHeader(http.StatusOK)
		jwt.Token = token
		utils.ResponseJSON(w, jwt)
	}
}

//TokenVerifyMiddleWare : Takes the handler function and returns the next handler function
func (c Controller) TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Inside the token verifier")
		var errorObject models.Error
		authHeader := r.Header.Get("Authorization") //returns map header
		// auth header has "bearer"
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) == 2 {
			authToken := bearerToken[1] //token value
			token, error := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				//to validate the token -- (SigningMethodHMAC) - checks only the signature for validation
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Error")
				}
				//we retunr the secret
				return []byte("secret"), nil
			})

			if error != nil {
				errorObject.Message = error.Error()
				utils.RespondError(w, http.StatusUnauthorized, errorObject)
				return
			}

			//Else the token is valid
			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				errorObject.Message = error.Error()
				utils.RespondError(w, http.StatusUnauthorized, errorObject)
				return
			}
		} else {
			errorObject.Message = "Invalid Token"
			utils.RespondError(w, http.StatusUnauthorized, errorObject)
			return
		}
	})
}
