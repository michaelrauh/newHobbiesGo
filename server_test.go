package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func TestMainFunction(t *testing.T) {
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
		if relativePath != "/newUser" {
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
