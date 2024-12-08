package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tolma-ch/tolmach-go/controllers/user"
	"github.com/tolma-ch/tolmach-go/routes"
)

func TestLogin(t *testing.T) {
	router := routes.MainRouter()
	authData := user.LoginInputData{
		Username: "foo",
		Password: "bar",
	}

	authDataJson, _ := json.Marshal(authData)

	w := httptest.NewRecorder()

	req, err := http.NewRequest("POST", "/login", strings.NewReader(string(authDataJson)))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)

	authData = user.LoginInputData{
		Username: "admin",
		Password: "password",
	}

	authDataJson, _ = json.Marshal(authData)

	w = httptest.NewRecorder()

	req, err = http.NewRequest("POST", "/login", strings.NewReader(string(authDataJson)))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}
