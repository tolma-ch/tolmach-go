package controller

import (
	"fmt"
	"net/http"

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
func GetTask(c *gin.Context, input *getTaskInput) (*Task, error) {
	for id, task := range taskStore {
		if id == input.Id {
			return &task, nil
		}
	}

	// Task not found
	c.AbortWithStatus(http.StatusNotFound)
	return nil, fmt.Errorf("task id=%d not found", input.Id)
}
