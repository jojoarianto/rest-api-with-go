# Rest API With Go
Simple Rest API with **Go** and **mongoDB**

The dependency of this REST API is : **github.com/gorilla/mux** mux library, **gopkg.in/mgo.v2** mongoDB driver,  **github.com/BurntSushi/toml** toml parser

## Installation & Run
install the dependency
```bash
# Download dependency of project
go get github.com/gorilla/mux
go get gopkg.in/mgo.v2
go github.com/BurntSushi/toml 
```
or if you have glide on your local you can use **glide** command
```bash
# Download dependency of project via glide
glide install
```
Build and run main.go 
```bash
go run main.go
```
## API

#### /users
* `GET` : Get all users
* `POST` : Create a new user
* `PUT` : Update an existing user 
* `DELETE` : Delete an existing user 

#### /users/:id
* `GET` : Get user by id

## Usage
Get All Data User

    curl -i -H "Accept: application/json" http://localhost:8000/users

Get spesific user by id

    curl -i -H "Accept: application/json" http://localhost:8000/users/5a799030d262cd04bd6fcac6

Insert data user

    curl -i -H "Content-Type: application/json" -X POST -d '{"name":"Irianto Jr","email":"irianto123@gmail.com", "no_hp":"0812345678", "event_id":"5a1b9bf6f8747703e8629414"}' http://localhost:8000/users

Update data user

    curl -i -H "Content-Type: application/json" -X PUT -d '{"id":"5ba17557ff74880e3b384a32","name":"EditiedBRO","email":"edites@gmail.com","no_hp":"0812345678","event_id":"5a1b9bf6f8747703e8629414"}' http://localhost:8000/users

Delete data user

	curl -i -H "Content-Type: application/json" -X DELETE -d '{"id":"5ba17557ff74880e3b384a32","name":"EditiedBRO","email":"edites@gmail.com","no_hp":"0812345678","event_id":"5a1b9bf6f8747703e8629414"}' http://localhost:8000/users	
