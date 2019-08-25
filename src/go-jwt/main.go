package main

import (
	"database/sql"
	"log"
	"net/http"

	"./controllers"
	"./driver"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

//global variable for DB connection
var db *sql.DB //will be available inside main and other functions

//Init func
func init() {
	gotenv.Load() //Loading the env variables from .env
}

//Main func
func main() {
	db = driver.ConnectDB()
	controller := controllers.Controller{}
	//Using gorilla-mux package to create the router for the application
	router := mux.NewRouter()
	router.HandleFunc("/signup", controller.Singup(db)).Methods("POST")
	router.HandleFunc("/login", controller.Login(db)).Methods("POST")
	router.HandleFunc("/protected", controller.TokenVerifyMiddleWare(controller.ProtectedEndpoint())).Methods("GET")

	log.Println("Server started successfully on port 8000.!!!")
	//http package -> to start the server
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatalf("Error while starting the application.")
	}
}
