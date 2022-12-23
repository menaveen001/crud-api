package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestInitRouter(t *testing.T) {
	router := InitRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "", w.Body.String())
}
