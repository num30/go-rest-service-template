package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/num30/go-rest-service-template/cmd/version"
	"github.com/num30/go-rest-service-template/pkg"
	"github.com/num30/go-rest-service-template/pkg/metrics"
	"github.com/num30/go-rest-service-template/pkg/rest"
	"github.com/sirupsen/logrus"
	"github.com/wI2L/fizz"
)

type Router struct {
	fizz                  *fizz.Fizz
	port                  int
	storageRequestTimeout time.Duration
}

func (r *Router) Run() {
	logrus.Info("Listening on port ", r.port)
	err := r.fizz.Engine().Run(fmt.Sprint(":", r.port))
	if err != nil {
		panic(err)
	}
}

func NewRouter(config *rest.HttpConfig, debugMode bool) *Router {
	f := rest.NewFizzRouter(config, pkg.ServiceName, version.Version, debugMode)

	r := &Router{
		fizz: f,
		port: config.Port,
	}
	r.init()
	return r
}

func (r *Router) init() {
	fooGroup := r.fizz.Group("/foos", "Foo list", "Foo operations")
	fooGroup.GET("/:name", nil, tonic.Handler(r.fooGetHandler, http.StatusOK))

}

func (r *Router) fooGetHandler(c *gin.Context, req *FooGetRequest) (*FooGetResponse, error) {
	// Record custom metric
	metrics.RecordSuccessfulRequestMetric()

	return &FooGetResponse{
		Result: "bar",
	}, nil
}
