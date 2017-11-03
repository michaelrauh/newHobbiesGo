package main

import (
	"os"
	"testing"
)

func TestThatStartingTheDatabaseDoesntResultInPanic(t *testing.T) {
	start()

	if db == nil {
		t.Fail()
	}
	stop()
	os.Remove("test.db")
}

func TestThatAddUserCreatesANewUser(t *testing.T) {
	start()
	var count int
	addUser("some_guid")
	db.Table("users").Count(&count)
	if count != 1 {
		t.Fail()
	}
	stop()
	os.Remove("test.db")
}
