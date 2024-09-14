package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("expected status code: %d, got %d", http.StatusOK, status)
	}

	body := responseRecorder.Body.String()
	if len(body) == 0 {
		t.Error("body is empty")
	}
}
func TestWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=1&city=unknown", nil)

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	if status != http.StatusBadRequest {
		require.Equal(t, http.StatusBadRequest, status, "wrong city value")
	} else {
		body := responseRecorder.Body.String()
		assert.Contains(t, body, "wrong city value", "body should contain wrong city value error")
	}
}
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	assert.Equal(t, http.StatusOK, status)

	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body)
	bodyList := strings.Split(body, ",")
	assert.Len(t, bodyList, totalCount)

}
