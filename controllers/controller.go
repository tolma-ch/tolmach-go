package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Task struct {
	Due     string `json:"due"`
	Remarks string `json:"remarks"`
}

// Mock task store
var taskStore = map[int]Task{
	0: {Due: "2022-01-01", Remarks: "this is very important!"},
	1: {Due: "2022-01-02", Remarks: "this is not important"},
}

type getTaskInput struct {
	Id int `path:"id"`
}

// Handler for /task/:id GET
func GetTask(c *gin.Context) {
	param, _ := strconv.Atoi(c.Param("id"))
	for id, task := range taskStore {
		if id == param {
			c.JSON(http.StatusOK, task)
		}
	}

	// Task not found
	c.AbortWithStatus(http.StatusNotFound)
}
