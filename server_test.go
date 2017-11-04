package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func TestMainFunction(t *testing.T) {

	calledAddRelated := false
	oldAddRelated := addRelated
	defer func() { addRelated = oldAddRelated }()

	addRelated = func(db *gorm.DB, in, rel string) {
		if rel != "text" || in != "other" {
			t.Fail()
		}
		calledAddRelated = true
	}

	calledAddH := false
	oldAddH := addH
	defer func() { addH = oldAddH }()

	addH = func(db *gorm.DB, in string) {
		if in != "text" {
			t.Fail()
		}
		calledAddH = true
	}

	var startedDB = false
	oldSt := st
	defer func() { st = oldSt }()

	st = func(kind, loc string) *gorm.DB {
		if kind != "sqlite3" || loc != "test.db" {
			t.Fail()
		}
		startedDB = true
		return nil
	}

	var got = false
	oldGet := get
	defer func() { get = oldGet }()

	get = func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
		if !(relativePath == "/newUser" || relativePath == "/hobbies") {
			t.Fail()
		}
		got = true
		return nil
	}

	var done = false
	oldRun := run
	defer func() { run = oldRun }()

	run = func(...string) error {
		done = true
		return nil
	}

	main()

	if !done {
		t.Fail()
	}

	if !got {
		t.Fail()
	}

	if !startedDB {
		t.Fail()
	}

	if !calledAddH {
		t.Fail()
	}

	if !calledAddRelated {
		t.Fail()
	}
}

func TestThatnewUserReturnsOK(t *testing.T) {
	oldAddUniq := addUniq
	defer func() { addUniq = oldAddUniq }()
	var called = false

	addUniq = func(db *gorm.DB) string {
		called = true
		return "abc123"
	}

	r := gin.Default()
	r.GET("/newUser", newUser)
	req, _ := http.NewRequest("GET", "/newUser", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Status code on GET of newUser was incorrect, got: %d, want: %d.", w.Code, 200)
	}

	if w.Body.String() != "{\"userID\":\"abc123\"}" {
		t.Errorf("body on GET of newUser was incorrect, got: %s, want: %s.", w.Body.String(), "{\"userID\":\"abc123\"}")
	}

	if !called {
		t.Fail()
	}
}

func TestThatHobbiesReturnsOK(t *testing.T) {
	oldAllH := allH
	defer func() { allH = oldAllH }()

	allH = func(db *gorm.DB) []hobby {
		hobbies := [2]hobby{hobby{Text: "this is a hobby"}}
		var s = hobbies[0:1]
		return s
	}

	r := gin.Default()
	r.GET("/hobbies", hobbies)
	req, _ := http.NewRequest("GET", "/hobbies", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Status code on GET of hobbies was incorrect, got: %d, want: %d.", w.Code, 200)
	}

	if !strings.Contains(w.Body.String(), "this is a hobby") {
		t.Fail()
	}
}
