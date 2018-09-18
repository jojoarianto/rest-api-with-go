package main

import (
	"fmt"
	// import package http for go web dev
	"net/http"
	// package for log print
	"log"
	// json library
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

// method handler for get all user
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, users)
}

// method handler for get user by id
func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // mux library to get all parameters
	user, err := dao.FindUserById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}
	respondWithJson(w, http.StatusOK, user)
}

// init the connection
// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	// router with gorilla mux initiate
	var r = mux.NewRouter()

	r.HandleFunc("/users", GetAllUsers).Methods("GET")      // call get method for get all user
	r.HandleFunc("/users/{id}", GetUserById).Methods("GET") // call get method for get spesific user by its id

	port := ":8000" // port for run the app

	fmt.Println("Start listening on port", port) // server up

	// Serve server on a port
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal() // print error log
	}

}

// method to print error output for http respon
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

// method to print output for http respon
// parameter  [w (Http.RestponWriter), http.statuscode, payload/data/msg]
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
