package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	router := UserRouter(gin.Default())
	authData := LoginInputData{
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

	authData = LoginInputData{
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
