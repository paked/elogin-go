package elogin

//PACKAGE DESCRIPTION
//	-Init: initialize connections
//	-Login(username, password): Login with details
//	-Register(username, password): Register user

import (
	// "errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type User struct {
	Username string
	Password string
}

var connection *mgo.Session
var con *mgo.Collection

var database string = "users-elogin-v1"

func Init() {
	session, err := mgo.Dial("localhost:27017")
	connection = session
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	c := connection.DB(database).C("users")
	err = c.Insert(&User{"harrison", "xkcd"}, &User{"xyz", "homie"})
	if err != nil {
		panic(err)
	}

	con = c

	result := User{}
	err = c.Find(bson.M{"username": "harrison"}).One(&result)
	// log.Println("password:", result.Password)
	if err != nil {
		panic(err)
	}

	log.Println("Hello v2")
}

func connectToDatabase(db string, collection string) (*mgo.Session, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	// defer session.Close()

	// c := connection.DB(db).C(collection)

	return session, err
}

func Login(username string, password string) (User, error) {
	sesh, err := connectToDatabase("", "") //mgo.Dial("localhost:27017")
	defer sesh.Close()
	c := sesh.DB(database).C("users")
	user := User{}
	err = c.Find(bson.M{"username": username, "password": password}).One(&user)

	return user, err
}

func Register() {

}
