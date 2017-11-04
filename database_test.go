package main

import (
	"os"
	"testing"
)

func TestThatStartingTheDatabaseDoesntResultInPanic(t *testing.T) {
	db = start("sqlite3", "test.db")

	if db == nil {
		t.Fail()
	}
	stop(db)
	os.Remove("test.db")
}

func TestThatAddUserCreatesANewUser(t *testing.T) {
	db = start("sqlite3", "test.db")
	var count int
	addUser(db, "some_guid")
	db.Table("users").Count(&count)
	if count != 1 {
		t.Fail()
	}
	stop(db)
	os.Remove("test.db")
}

func TestThatExistsReturnsTrueWhenGUIDIsUsedAndFalseOtherwise(t *testing.T) {
	var found bool
	db = start("sqlite3", "test.db")

	found = exist(db, "some_guid")

	if found {
		t.Fail()
	}

	addUser(db, "some_guid")

	found = exist(db, "some_guid")

	if !found {
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

func TestThatAddHobbyCreatesANewHobby(t *testing.T) {
	db = start("sqlite3", "test.db")
	addHobby(db, "some description")
	var hobbies []hobby
	db.Find(&hobbies)
	if len(hobbies) != 1 || hobbies[0].Text != "some description" {
		t.Fail()
	}

	stop(db)
	os.Remove("test.db")
}

func TestThatAllHobbiesGetsAllHobbies(t *testing.T) {
	db = start("sqlite3", "test.db")

	addHobby(db, "some description")
	addHobby(db, "other description")

	res := allHobbies(db)
	if len(res) != 2 || res[1].Text != "other description" {
		t.Fail()
	}
	stop(db)
	os.Remove("test.db")
}

func TestThatAddRelatedHobbyFindsAReferenceToOtherHobby(t *testing.T) {
	db = start("sqlite3", "test.db")

	addHobby(db, "some description")
	addRelatedHobby(db, "other description", "some description")

	if allHobbies(db)[1].Prerequisite != 1 {
		t.Fail()
	}

	stop(db)
	os.Remove("test.db")
}
