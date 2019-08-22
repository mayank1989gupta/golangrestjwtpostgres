package main

import "net/http"

//Method to handle the login request
func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Successfully called logged!!"))
}
