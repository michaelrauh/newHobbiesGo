package main

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var err error

func Start() *gorm.DB {
  db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&user{})

  return db
}

func Stop(db *gorm.DB) {
  db.Close()
}

func AddUser(db *gorm.DB, guid string) {
  db.Create(&user{GUID: guid})
}

type user struct {
	gorm.Model
	GUID string
}
