package main

import (
	"os"
	"testing"
)

func TestThatStartingTheDatabaseDoesntResultInPanic(t *testing.T) {
	db := Start()

	if db == nil {
		t.Fail()
	}
	Stop(db)
	os.Remove("test.db")
}

func TestThatAddUserCreatesANewUser(t *testing.T) {
	db := Start()
	var count int
	AddUser(db, "some_guid")
	db.Table("users").Count(&count)
	if count != 1 {
		t.Fail()
	}
	Stop(db)
	os.Remove("test.db")
}
