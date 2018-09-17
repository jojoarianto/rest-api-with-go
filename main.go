package main

import (
	"fmt"
	// import package http for go web dev
	"net/http"
	// package for log print
	"log"

	// gorila mux, router n dispatcher fo go
	"github.com/gorilla/mux"
	// import all local project package
	. "github.com/jojoarianto/rest-api-with-go/config" // construct read to read db config with toml
	. "github.com/jojoarianto/rest-api-with-go/dao"    // data access object
	. "github.com/jojoarianto/rest-api-with-go/models" // to import all model (include User)
)

// method for get all user
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, users)

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
