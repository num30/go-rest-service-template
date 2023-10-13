package main

import (
	"log/slog"
	"os"

	"github.com/num30/config"
	"github.com/num30/go-rest-service-template/pkg"
	"github.com/num30/go-rest-service-template/pkg/metrics"

	"github.com/num30/go-rest-service-template/cmd/version"

	ccmd "github.com/num30/go-rest-service-template/pkg/cmd"
	"github.com/num30/go-rest-service-template/pkg/server"
)

var serviceConfig = &pkg.Config{}

func main() {
	ccmd.ProcessVersionArgument(pkg.ServiceName, os.Args, version.Version)

	loadConfig()

	metrics.StartMetricsServer(&serviceConfig.Metrics)
	r := server.NewRouter(&serviceConfig.Service, serviceConfig.IsDebugMode)
	r.Run()
}

// loadConfig reads in config file, ENV variables, and flags if set.
func loadConfig() {
	err := config.NewConfReader("service_test").Read(serviceConfig)
	if err != nil {
		slog.With("error", err).Error("Error reading config")
		os.Exit(1)
	}
}
