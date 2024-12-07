package user

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	mw "github.com/tolma-ch/tolmach-go/middleware"
)

type LoginInputData struct {
	Username string `example:"username"`
	Password string `example:"pass"`
}

type LoginOutputData struct {
	Token string `example:"gadsgasfdasfds"`
}

// @Summary Login user
// @Schemes
// @Description Login user based on provided `username` and `password`
// @Tags User
// @Accept json
// @Produce json
// @Param form body LoginInputData true "name search by q"
// @Success 200 {object} LoginOutputData
// @Router /login [post]
func Login(c *gin.Context) {
	// In a real application, authenticate the user (this is just an example)
	data := LoginInputData{}

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "fckff"})
		return
	}

	// Check user credentials
	var role string
	if data.Username == "admin" && data.Password == "password" {
		role = "admin"
	} else if data.Username == "user" && data.Password == "password" {
		role = "user"
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": data.Username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expiration time
	})

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(mw.SecretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
	}

	c.JSON(http.StatusOK, LoginOutputData{Token: tokenString})
}

func handleAdminResource(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Admin resource accessed successfully",
	})
}
