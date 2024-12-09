package user

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tolma-ch/tolmach-go/helpers"
)

func TestLogin(t *testing.T) {
	router := UserRouter(gin.Default())
	authData := LoginInputData{
		Username: "foo",
		Password: "bar",
	}

	return_code := helpers.JsonPOSTTest(router, "/login", authData)

	assert.Equal(t, 401, return_code)

	authData = LoginInputData{
		Username: "admin",
		Password: "password",
	}

	return_code = helpers.JsonPOSTTest(router, "/login", authData)
	assert.Equal(t, 200, return_code)

}
