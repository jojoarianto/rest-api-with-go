// Data Access Object
package dao

import (
	"log" // to print error log

	. "github.com/jojoarianto/rest-api-go/models" // to import all model (include User)
	mgo "gopkg.in/mgo.v2"                         // mongo db driver package
	"gopkg.in/mgo.v2/bson"                        // bson for data structure
)

type UsersDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "users" // Colection name
)

// implies, establish a connection to MongoDB database
func (m *UsersDAO) Connect() {
	session, err := mgo.Dial(m.Server) // establish connection
	if err != nil {                    // if connection failed to establish
		log.Fatal(err) // print error log
	}

	db = session.DB(m.Database)
}

// queries method for get all data user
func (m *UsersDAO) GetAll() ([]User, err) {
	var users []User
	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}
