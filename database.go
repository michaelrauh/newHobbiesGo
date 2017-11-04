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
	db.AutoMigrate(&hobby{})

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

func addHobby(db *gorm.DB, text string) {
	db.Create(&hobby{Text: text})
}

func allHobbies(db *gorm.DB) []hobby {
	var hobbies []hobby
	db.Find(&hobbies)
	return hobbies
}

type user struct {
	gorm.Model
	GUID string
}

type hobby struct {
	gorm.Model
	Text string
}
