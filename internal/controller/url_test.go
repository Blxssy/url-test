package controller

import (
	"github.com/Blxssy/url-test/internal/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUrlController_CreateShortURL(t *testing.T) {
	router, container := test.PrepareForControllerTest()

	url := NewURLController(container)
	router.POST("/", url.CreateShortURL)

	body := `{"url": "https://example.com"}`
	req, _ := http.NewRequest("POST", "/", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
