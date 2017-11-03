package main

import (
	"os"
	"testing"
)

func TestThatStartingTheDatabaseDoesntResultInPanic(t *testing.T) {
	db := start()

	if db == nil {
		t.Fail()
	}
	stop(db)
	os.Remove("test.db")
}

func TestThatAddUserCreatesANewUser(t *testing.T) {
	db := start()
	var count int
	addUser(db, "some_guid")
	db.Table("users").Count(&count)
	if count != 1 {
		t.Fail()
	}
	stop(db)
	os.Remove("test.db")
}
