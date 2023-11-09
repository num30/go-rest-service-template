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
	slog.Info("Listening", "port", r.port)
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

	// Define routes here
	thingsGroup := r.fizz.Group("/things", "Things list", "Things operations")

	thingsGroup.PUT("/:name",
		[]fizz.OperationOption{fizz.Summary("Put a thing")},
		tonic.Handler(r.thingPutHandler, http.StatusCreated))

	thingsGroup.GET("/:name", []fizz.OperationOption{fizz.Summary("Get a thing")}, tonic.Handler(r.thingGetHandler, http.StatusOK))
}

func (r *Router) thingGetHandler(c *gin.Context, req *ThingGetRequest) (*ThingGetResponse, error) {
	// Record custom metric
	metrics.RecordSuccessfulRequestMetric()
	if val, ok := r.testStorage[req.Name]; ok {
		return &ThingGetResponse{
			Result: val,
		}, nil
	} else {
		return nil, rest.HttpError{
			HttpCode: http.StatusNotFound,
			Message:  "Not found",
		}
	}
}

func (r *Router) thingPutHandler(c *gin.Context, req *ThingPutRequest) (*ThingPutResponse, error) {
	r.testStorage[req.Name] = req.Value
	return &ThingPutResponse{}, nil
}
