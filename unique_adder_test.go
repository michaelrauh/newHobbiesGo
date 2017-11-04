package main

import (
	"testing"

	"github.com/jinzhu/gorm"
)

func TestThatAddUniqueWillAttemptToAddUntilSuccessful(t *testing.T) {
	oldSafeAdder := safeAdder
	defer func() { safeAdder = oldSafeAdder }()
	var iter = 0
	safeAdder = func(db *gorm.DB) (bool, string) {
		iter++
		if iter < 5 {
			return false, "no"
		}
		return true, "yes"
	}

	var res = addUnique(db)
	if iter < 5 || res != "yes" {
		t.Fail()
	}
}
