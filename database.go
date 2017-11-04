package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func start(database, name string) *gorm.DB {
	db, err := gorm.Open(database, name)
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&user{})

	return db
}

func exist(db *gorm.DB, guid string) bool {
	var u user
	db.Where("GUID = ?", guid).First(&u)
	return u.ID != 0
}

func stop(db *gorm.DB) {
	db.Close()
}

func addUser(db *gorm.DB, guid string) {
	db.Create(&user{GUID: guid})
}

type user struct {
	gorm.Model
	GUID string
}
