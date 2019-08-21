package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
)

//global variable for DB connection
var db *sql.DB //will be available inside main and other functions

func main() {
	pgURL, err := pq.ParseURL("postgres://xzongdeg:qffSCfAxoD_u3k1IUCwERreKeNtKtpn0@isilo.db.elephantsql.com:5432/xzongdeg")
	if err != nil {
		log.Fatalln("Error while connecting to DB")
		os.Exit(1)
	}

	//to open a connection -> db is the global variable
	db, err = sql.Open("postgres", pgURL)
	if err != nil {
		log.Fatal(err)
	}

	// ping db to test
	err = db.Ping() //if ping does not retunr anything that means connection is established successfully
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	//Using gorilla-mux package to create the router for the application
	router := mux.NewRouter()
	router.HandleFunc("/signup", singup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/protected", TokenVerifyMiddleWare(protectedEndpoint)).Methods("GET")

	log.Println("Server started successfully on port 8000.!!!")
	//http package -> to start the server
	err = http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatalf("Error while starting the application.")
	}
}

//Method to handle the singup request
func singup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Successfully called signed!!"))
}

//Method to handle the login request
func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Successfully called logged!!"))
}

//Method for protectedEndpoint -> would accessed by jwt token
func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("protectedEndpoint Invoked!")
}

//TokenVerifyMiddleWare : Takes the handler function and returns the next handler function
func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("/TokenVerifyMiddleWare Invoked!")
	return nil
}
