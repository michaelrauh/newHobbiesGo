package main

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func start()  {
  db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&user{})
}

func stop() {
  db.Close()
}

func addUser(guid string) {
  db.Create(&user{GUID: guid})
}

type user struct {
	gorm.Model
	GUID string
}
