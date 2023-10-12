//go:build stest
// +build stest

package service

import (
	"fmt"
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
}

var tConfig = &TestConfig{}

func init() {
	err := config.NewConfReader("service_test").Read(tConfig)
	if err != nil {
		slog.With("error", err).Error("Error reading config")
	}
}

// Call endpoint and check response
func Test_AddCApp(t *testing.T) {
	rest := resty.New()
	t.Run("Ping", func(t *testing.T) {
		r, err := rest.R().Get(tConfig.ServiceHost + "/ping")
		assert.NoError(t, err)
		assert.Equal(t, r.StatusCode(), http.StatusOK)
	})

	t.Run("Get", func(tt *testing.T) {
		r, err := rest.R().Get(fmt.Sprintf("%s/%s", tConfig.ServiceHost, "some-name"))
		if err != nil {
			assert.FailNow(tt, err.Error())
		}
		assert.Equal(tt, r.StatusCode(), http.StatusOK)
		assert.NoError(tt, err)
	})
}
