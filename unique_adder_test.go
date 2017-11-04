package main

import (
	"testing"

	"github.com/jinzhu/gorm"
)

func TestThatAddUniqueWillAttemptToAddUntilSuccessful(t *testing.T) {
	oldSafeAdder := safeAdder
	defer func() { safeAdder = oldSafeAdder }()
	var iter = 0
	safeAdder = func(db *gorm.DB) bool {
		iter++

		if iter < 5 {
			return false
		}
		return true
	}

	addUnique(db)
	if iter < 5 {
		t.Fail()
	}
}
