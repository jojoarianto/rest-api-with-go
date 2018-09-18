// Data Access Object
package dao

import (
	"log" // to print error log

	. "github.com/jojoarianto/rest-api-with-go/models" // to import all model (include User)
	mgo "gopkg.in/mgo.v2"                              // mongo db driver package
	"gopkg.in/mgo.v2/bson"                             // bson for data structure
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

// queries for get all data user
func (m *UsersDAO) GetAll() ([]User, error) {
	var users []User
	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

// query for get user by their id
func (m *UsersDAO) FindUserById(id string) (User, error) {
	var user User
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user) // get one data by id
	return user, err
}

// query to post user
func (m *UsersDAO) Insert(user User) error {
	err := db.C(COLLECTION).Insert(&user)
	return err
}

// query to update data user
func (m *UsersDAO) Update(user User) error {
	err := db.C(COLLECTION).UpdateId(user.ID, &user)
	return err
}

// query for delete user
func (m *UsersDAO) Delete(user User) error {
	err := db.C(COLLECTION).Remove(user)
	return err
}
