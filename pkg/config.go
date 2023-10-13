package pkg

import (
	"github.com/num30/go-rest-service-template/pkg/metrics"
	"github.com/num30/go-rest-service-template/pkg/rest"
)

type Config struct {
	IsDebugMode bool   `default:"false" envvar:"DEBUG_MODE"`
	LogLevel    string `default:"info" envvar:"LOG_LEVEL"`
	Metrics     metrics.MetricsConfig
	Service     rest.HttpConfig
}
