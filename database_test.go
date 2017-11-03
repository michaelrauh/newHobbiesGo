package main

import (
	"os"
	"testing"
)

func TestThatStartingTheDatabaseDoesntResultInPanic(t *testing.T) {
	db := start("sqlite3", "test.db")

	if db == nil {
		t.Fail()
	}
	stop(db)
	os.Remove("test.db")
}

func TestThatAddUserCreatesANewUser(t *testing.T) {
	db := start("sqlite3", "test.db")
	var count int
	addUser(db, "some_guid")
	db.Table("users").Count(&count)
	if count != 1 {
		t.Fail()
	}
	stop(db)
	os.Remove("test.db")
}

func TestThatStartCanPanic(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			_, ok := r.(error)
			if ok {
				t.Fail()
			}
		}

	}()

	start("foolite", "nope")

}
