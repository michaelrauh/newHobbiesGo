package main

import (
	"github.com/jinzhu/gorm"
)

var gen = Generate
var add = addUser
var ex = exist

func tryToAdd(db *gorm.DB) (bool, string) {
	guid := gen(6)
	existsAlready := ex(db, guid)
	if !existsAlready {
		add(db, guid)
	}
	return !existsAlready, guid
}
