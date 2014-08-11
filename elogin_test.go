package elogin

import (
	"testing"
)

const testUsername string = "bob"
const testPassword string = "cherryTreeLane"

func TestRegister(t *testing.T) {
	Init()

	user, err := Register(testUsername, testPassword)

	if err != nil {
		t.Errorf("mgo database error")
	}

	if user == (User{}) {
		TestRemove(t)
		t.Errorf("A user with that name already exists")
	}
}

func TestLogin(t *testing.T) {
	Init()

	user, err := Login(testUsername, testPassword)
	if user == (User{}) {
		t.Errorf("A user with that username and password combination does not exist")
		t.FailNow()
	}

	if err != nil {
		t.Errorf("mgo database error")
		t.FailNow()
	}
}

func TestRemove(t *testing.T) {
	Init()

	err := Remove(testUsername, testPassword)
	if err != nil {
		t.Error("Could not delete user. DB user")
	}
}
