package main

import (
	"testing"

	"github.com/jinzhu/gorm"
)

func TestThatTryToAddWillAddAGUID(t *testing.T) {
	oldGen := gen
	defer func() { gen = oldGen }()

	gen = func(in int) string {
		if in != 6 {
			t.Errorf("length of id not what expected, got: %d, want: %d.", in, 6)
		}
		return "abc123"
	}

	oldAdd := add
	defer func() { add = oldAdd }()
	var called = false

	add = func(db *gorm.DB, guid string) {
		if guid != "abc123" {
			t.Fail()
		}
		called = true
	}

	oldEx := ex
	defer func() { ex = oldEx }()

	ex = func(db *gorm.DB, guid string) bool {
		return false
	}

	ok, res := tryToAdd(db)

	if !called || !ok {
		t.Fail()
	}

	if len(res) < 6 {
		t.Fail()
	}
}

func TestThatTryToAddWillNotAddAUsedGUID(t *testing.T) {
	oldAdd := add
	defer func() { add = oldAdd }()
	var called = false

	add = func(db *gorm.DB, guid string) {
		called = true
	}

	oldEx := ex
	defer func() { ex = oldEx }()

	ex = func(db *gorm.DB, guid string) bool {
		return true
	}

	ok, _ := tryToAdd(db)

	if called || ok {
		t.Fail()
	}
}
