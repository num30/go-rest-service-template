//go:build servicetest
// +build servicetest

package service

import (
	"log/slog"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/num30/config"
	"github.com/stretchr/testify/assert"
)

type TestConfig struct {
	LogLevel    string `default:"info" envvar:"LOG_LEVEL"`
	ServiceHost string `default:"http://localhost:8080" envvar:"SERVICE_HOST"`
	MetricsHost string `default:"http://localhost:10250" envvar:"METRICS_HOST"`
}

var tConfig = &TestConfig{}
var restClient = resty.New()

func init() {
	err := config.NewConfReader("service_test").Read(tConfig)
	if err != nil {
		slog.With("error", err).Error("Error reading config")
	}
}

func Test_Ping(t *testing.T) {
	t.Run("Ping", func(t *testing.T) {
		r, err := restClient.R().Get(tConfig.ServiceHost + "/ping")
		assert.NoError(t, err)
		assert.Equal(t, r.StatusCode(), http.StatusOK)
	})
}

// Call endpoint and check response
func Test_Things(t *testing.T) {
	// call the PUT endpoint
	t.Run("Add", func(tt *testing.T) {
		r, err := restClient.R().SetBody(`{ "value": "Thing value" }`).Put(tConfig.ServiceHost + "/things/some-name")
		if assert.NoError(tt, err) {
			assert.Equal(tt, http.StatusCreated, r.StatusCode(), "Invalid status code. Response body: %s", string(r.Body()))
		}
	})

	// call the GET endpoint
	t.Run("Get", func(tt *testing.T) {
		r, err := restClient.R().Get(tConfig.ServiceHost + "/things/some-name")
		if assert.NoError(tt, err) {
			assert.Equal(tt, http.StatusOK, r.StatusCode(), "Invalid status code. Response body: %s", string(r.Body()))
		}
	})
}

// Check that prometheus metrics are available
func Test_GetMetrics(t *testing.T) {
	t.Run("Metrics", func(t *testing.T) {
		r, err := restClient.R().Get(tConfig.MetricsHost + "/metrics")
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, r.StatusCode())
			assert.Contains(t, string(r.Body()), "things_requests_total")
		}
	})
}
