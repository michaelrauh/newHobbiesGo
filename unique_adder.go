package main

import (
	"github.com/jinzhu/gorm"
)

var safeAdder = tryToAdd

func addUnique(db *gorm.DB) string {
	var res string
	for ok := false; !ok; ok, res = safeAdder(db) {
	}
	return res
}
