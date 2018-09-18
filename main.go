package main

import (
	"fmt"
	// import package http for go web dev
	"net/http"
	// package for log print
	"log"

	"encoding/json"

	// gorila mux, router n dispatcher fo go
	"github.com/gorilla/mux"
	// import all local project package
	. "github.com/jojoarianto/rest-api-with-go/config" // construct read to read db config with toml
	. "github.com/jojoarianto/rest-api-with-go/dao"    // data access object
	// . "github.com/jojoarianto/rest-api-with-go/models" // to import all model (include User)
)

var config = Config{}
var dao = UsersDAO{}

// method for get all user
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, users)

}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
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
