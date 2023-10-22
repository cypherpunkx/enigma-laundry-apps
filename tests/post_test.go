package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAllPosts(t *testing.T) {

	router := gin.Default()

	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1", nil)

	if err != nil {
		t.Errorf("expected bla but got 1 %s", err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, `{"message":"Hello, World!"}`, rr.Body.String())
}
