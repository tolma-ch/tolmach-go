package helpers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
)

func JsonPOSTTest(router *gin.Engine, url string, data any) int {
	w := httptest.NewRecorder()

	jsonData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(jsonData)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	return w.Code
}
