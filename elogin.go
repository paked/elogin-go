package elogin

//PACKAGE DESCRIPTION
//	-Init: initialize connections
//	-Login(username, password): Login with details
//	-Register(username, password): Register user
// 	-Remove(username, password): Remove a user document

import (
	"code.google.com/p/go.crypto/bcrypt"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Username  string
	Password  string
	Timestamp time.Time
}

var connection *mgo.Session
var con *mgo.Collection

const database string = "users-elogin-v1"

const userDB string = "users"

func Init() {
	session, err := mgo.Dial("localhost:27017")
	connection = session
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	// log.Println("password:", result.Password)

	c := session.DB(database).C(userDB)
	//
	user := []User{}
	err = c.Find(bson.M{}).Sort("timestamp").All(&user)
	fmt.Printf("")

	// c.RemoveAll(&User{})

}

func connectToDatabase() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost:27017")
	return session, err
}

func Login(username string, password string) (User, error) {
	sesh, err := connectToDatabase()
	if err != nil {
		panic(err)
	}
	defer sesh.Close()

	c := sesh.DB(database).C(userDB)
	user := User{}
	err = c.Find(bson.M{"username": username, "password": password}).One(&user)

	sesh.Close()
	return user, err
}

func Register(username string, password string) (User, error) {
	sesh, err := connectToDatabase()
	if err != nil {
		panic(err)
	}
	defer sesh.Close()

	c := sesh.DB(database).C(userDB)
	user := User{}
	err = c.Find(bson.M{"username": username}).One(&user)
	passwordCrypt, err := Crypt([]byte(password))
	if user == (User{}) {
		newUser := User{username, string(passwordCrypt), time.Now()}
		c.Insert(&newUser)

		return newUser, nil
	}

	sesh.Close()

	return User{}, nil
}

func Remove(username string, password string) error {
	sesh, err := connectToDatabase()
	if err != nil {
		panic(err)
	}
	defer sesh.Close()

	c := sesh.DB(database).C(userDB)
	err = c.Remove(bson.M{"username": username, "password": password})
	return nil
}

func Crypt(password []byte) ([]byte, error) {
	defer clear(password)
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

func clear(b []byte) {
	for i := 0; i < len(b); i++ {
		b[i] = 0
	}
}
