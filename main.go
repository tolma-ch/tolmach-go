package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
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
func getTask(c *gin.Context, input *getTaskInput) (*Task, error) {
	for id, task := range taskStore {
		if id == input.Id {
			return &task, nil
		}
	}

	// Task not found
	c.AbortWithStatus(http.StatusNotFound)
	return nil, fmt.Errorf("task id=%d not found", input.Id)
}

func main() {
	router := gin.Default()
	f := fizz.NewFromEngine(router)
	f.GET("/task/:id", nil, tonic.Handler(getTask, http.StatusOK))

	info := &openapi.Info{
		Title:       "Task API",
		Description: "manage tasks",
		Version:     "0.0.1",
	}
	f.GET("/openapi.json", nil, f.OpenAPI(info, "json"))

	srv := &http.Server{
		Addr:    "localhost:8000",
		Handler: f,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

	router.Run("localhost:8000")
}
