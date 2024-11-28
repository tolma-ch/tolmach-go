package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tolma-ch/tolmach-go/routes"
)

func TestLogin(t *testing.T) {
	router := routes.MainRouter()
	data := url.Values{}
	data.Set("username", "foo")
	data.Set("password", "bar")

	w := httptest.NewRecorder()

	// userJson, _ := json.Marshal(map[string]string{"username": "", "password": ""})
	req, err := http.NewRequest("POST", "/login", strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println(err)
	}
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)

	// correctUserJson, _ := json.Marshal(map[string]string{"username": "admin", "password": "password"})
	data.Set("username", "admin")
	data.Set("password", "password")
	w = httptest.NewRecorder()
	req, err = http.NewRequest("POST", "/login", strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}
