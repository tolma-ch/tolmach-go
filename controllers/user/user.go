package user

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	mw "github.com/tolma-ch/tolmach-go/middleware"
)

func Login(c *gin.Context) {
	// In a real application, authenticate the user (this is just an example)
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Check user credentials
	var role string
	if username == "admin" && password == "password" {
		role = "admin"
	} else if username == "user" && password == "password" {
		role = "user"
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expiration time
	})

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(mw.SecretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func handleAdminResource(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Admin resource accessed successfully",
	})
}
