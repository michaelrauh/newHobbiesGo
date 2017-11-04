package main

import (
	"github.com/jinzhu/gorm"
)

var safeAdder = tryToAdd

func addUnique(db *gorm.DB) {
	for !safeAdder(db) {
	}
}
