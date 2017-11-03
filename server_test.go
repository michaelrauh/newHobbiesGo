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
	if db == nil {
		t.Fail()
	}

	if !done {
		t.Fail()
	}

	if !got {
		t.Fail()
	}
}

func TestThatnewUserReturnsOK(t *testing.T) {
	oldGen := gen
	defer func() { gen = oldGen }()

	gen = func(in int) string {
		if in != 6 {
			t.Errorf("length of id not what expected, got: %d, want: %d.", in, 6)
		}
		return "abc123"
	}

	oldAdd := add
	defer func() { add = oldAdd }()
	var called = false

	add = func(db *gorm.DB, guid string) {
		called = true
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
