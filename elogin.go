package elogin

//PACKAGE DESCRIPTION
//	-Init: initialize connections
//	-Login(username, password): Login with details
//	-Register(username, password): Register user
// 	-Remove(username, password): Remove a user document

import (
	"code.google.com/p/go.crypto/bcrypt"
	// "errors"
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

type Settings struct {
	URL        string `default: "localhost:27017"`
	Database   string `default: "users-elogin-v1"`
	Collection string `default: "users"`
}

var connection *mgo.Session
var con *mgo.Collection

var config Settings

func Init(settings Settings) {
	config = settings
	session, err := connectToDatabase()
	if err != nil {
		panic(err)
	}

	c := session.DB(config.Database).C(config.Collection)

	user := []User{}
	err = c.Find(bson.M{}).Sort("timestamp").All(&user)
	// fmt.Printf("%v", &user)/

	// c.RemoveAll(bson.M{})

}

func connectToDatabase() (*mgo.Session, error) {
	session, err := mgo.Dial(config.URL)
	return session, err
}

func Login(username string, password string) (User, error) {
	sesh, err := connectToDatabase()
	if err != nil {
		panic(err)
	}
	defer sesh.Close()

	c := sesh.DB(config.Database).C(config.Collection)
	user := User{}
	// err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	err = c.Find(bson.M{"username": username}).One(&user)
	if user != (User{}) {
		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil {
			return user, nil
		} else {
			return User{}, nil
			fmt.Printf("PASSWORDS DO NOT MATCH")
		}
	}
	sesh.Close()
	return user, nil
}

func Register(username string, password string) (User, error) {
	sesh, err := connectToDatabase()
	if err != nil {
		panic(err)
	}
	defer sesh.Close()

	c := sesh.DB(config.Database).C(config.Collection)
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

	c := sesh.DB(config.Database).C(config.Collection)
	err = c.Remove(bson.M{"username": username, "password": password})
	return nil
}

func Clean() error {
	sesh, err := connectToDatabase()
	if err != nil {
		panic(err)
	}
	defer sesh.Close()

	c := sesh.DB(config.Database).C(config.Collection)
	c.RemoveAll(bson.M{})
	return err
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
