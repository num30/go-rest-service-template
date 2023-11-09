package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
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
	r := NewRouter(defaultConfig, false).fizz.Engine()
	t.Run("PUT", func(tt *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/things/bar", strings.NewReader(`{ "value": "Thing value" }`))
		// act
		r.ServeHTTP(w, req)

		// assert
		assert.Equal(t, 201, w.Code)
	})

	t.Run("GET", func(tt *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/things/bar", nil)

		// act
		r.ServeHTTP(w, req)

		// assert
		assert.Equal(t, 200, w.Code)
	})

}
