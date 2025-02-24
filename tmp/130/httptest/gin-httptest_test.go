package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUser_Success(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/users/123", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"id":"123", "name":"John Doe"}`, w.Body.String())
}
