package main

import (
	"fmt"
	// import package http for go web dev
	"net/http"
	// package for log print
	"log"

	// gorila mux, router n dispatcher fo go
	"github.com/gorilla/mux"
)

// method for get all user
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not Implement!")
}

func main() {
	var r = mux.NewRouter()

	r.HandleFunc("/users", GetAllUsers).Methods("GET") // call get method for get all user

	port := ":8000" // port for run the app

	fmt.Println("Start listening on port", port) // server up

	// Serve server on a port
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal() // print error log
	}

}
