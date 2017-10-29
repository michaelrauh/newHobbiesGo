package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"os"
)


func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func TestThatnewUserReturnsOK(t *testing.T) {
	r := gin.Default()
	r.GET("/newUser", newUser)
	req, _ := http.NewRequest("GET", "/newUser", nil)
	w := httptest.NewRecorder()
  
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Status code on GET of newUser was incorrect, got: %d, want: %d.", w.Code, 200)
	}

	if w.Body.String() != "{\"userID\":\"abc123\"}" {
		t.Errorf("body on GET of newUser was incorrect, got: %d, want: %d.", w.Body.String(), "{\"userID\":\"abc123\"}")
	}
}
