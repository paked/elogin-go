package elogin

import (
	"testing"
)

const testUsername string = "bob"
const testPassword string = "cherryTreeLane"

var testConfig Settings = Settings{"localhost:27017", "users-elogin-v1", "test-users"}

var elogin Elogin

func TestInit(t *testing.T) {
	elogin = Elogin{}
	elogin.Init(testConfig)
}

func TestClean(t *testing.T) {
	err := elogin.Clean()
	if err != nil {
		t.Error("Something went wrong cleaning the test DB, rip your elegant statements.")
	}
}

func TestRegister(t *testing.T) {
	user, err := elogin.Register(testUsername, testPassword)

	if err != nil {
		t.Errorf("mgo database error")
	}

	if user == (User{}) {
		// TestRemove(t)
		t.Errorf("A user with that name already exists")
	}
}

func TestLoginSuccess(t *testing.T) {

	user, err := elogin.Login(testUsername, testPassword)
	if user == (User{}) {
		t.Errorf("A user with that username and password combination does not exist")
		t.FailNow()
	}

	if err != nil {
		t.Errorf("mgo database error")
		t.FailNow()
	}
}

func TestLoginPasswordFail(t *testing.T) {
	user, err := elogin.Login(testUsername, testPassword+"fail")
	if user != (User{}) {
		t.Errorf("Password checking broken")
	}

	if err != nil {
		t.Error("mgo database error")
	}
}

func TestLoginUsernameFail(t *testing.T) {
	user, err := elogin.Login(testUsername+"fail", testPassword)
	if user != (User{}) {
		t.Error("Username checking wrong")
	}

	if err != nil {
		t.Error(err)
	}
}

func TestRemove(t *testing.T) {
	err := elogin.Remove(testUsername, testPassword)
	if err != nil {
		t.Error("Could not delete user. DB user")
	}
}
