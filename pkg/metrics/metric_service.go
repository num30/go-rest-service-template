package metrics

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricsConfig struct {
	Port    int  `default:"10250" envvar:"METRICS_PORT"`
	Enabled bool `default:"true" envvar:"METRICS_ENABLED"`
}

func StartMetricsServer(conf *MetricsConfig) {
	if conf.Enabled {
		go func() {
			http.Handle("/metrics", promhttp.Handler())
			slog.Info("Starting metrics server", "port", conf.Port)
			err := http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil)
			if err != nil {
				slog.With("error", err).Error("Error starting metrics server")
			}
		}()

	}
}
