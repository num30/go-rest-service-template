package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/num30/go-rest-service-template/cmd/version"
	"github.com/num30/go-rest-service-template/pkg"
	"github.com/num30/go-rest-service-template/pkg/metrics"
	"github.com/num30/go-rest-service-template/pkg/rest"

	"github.com/wI2L/fizz"
)

type Router struct {
	fizz        *fizz.Fizz
	port        int
	testStorage map[string]string
}

func (r *Router) Run() {
	slog.Info("Listening on port ", r.port)
	err := r.fizz.Engine().Run(fmt.Sprint(":", r.port))
	if err != nil {
		panic(err)
	}
}

func NewRouter(config *rest.HttpConfig, debugMode bool) *Router {
	f := rest.NewFizzRouter(config, pkg.ServiceName, version.Version, debugMode)

	r := &Router{
		fizz:        f,
		port:        config.Port,
		testStorage: map[string]string{},
	}
	r.init()
	return r
}

func (r *Router) init() {
	fooGroup := r.fizz.Group("/foos", "Foo list", "Foo operations")
	fooGroup.PUT("/:name", nil, tonic.Handler(r.fooPutHandler, http.StatusCreated))
	fooGroup.GET("/:name", nil, tonic.Handler(r.fooGetHandler, http.StatusOK))
}

func (r *Router) fooGetHandler(c *gin.Context, req *FooGetRequest) (*FooGetResponse, error) {
	// Record custom metric
	metrics.RecordSuccessfulRequestMetric()
	if val, ok := r.testStorage[req.Name]; ok {
		return &FooGetResponse{
			Result: val,
		}, nil
	} else {
		return nil, rest.HttpError{
			HttpCode: http.StatusNotFound,
			Message:  "Not found",
		}
	}
}

func (r *Router) fooPutHandler(c *gin.Context, req *FooPutRequest) (*FooPutResponse, error) {
	r.testStorage[req.Name] = req.Value
	return &FooPutResponse{}, nil
}
