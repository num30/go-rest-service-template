package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/cors"
	"github.com/num30/go-rest-service-template/pkg/rest"
	"github.com/stretchr/testify/assert"
)

func init() {
}

var defaultConfig = &rest.HttpConfig{
	Cors: cors.Config{
		AllowAllOrigins: true,
	},
}

func Test_RouterPing(t *testing.T) {
	// arrange
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)

	r := NewRouter(defaultConfig, false).fizz.Engine()

	// act
	r.ServeHTTP(w, req)

	// assert
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "pong")
}

func Test_RouterGet(t *testing.T) {
	// arrange
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/foos/bar", nil)

	r := NewRouter(defaultConfig, false).fizz.Engine()

	// act
	r.ServeHTTP(w, req)

	// assert
	assert.Equal(t, 200, w.Code)
}
