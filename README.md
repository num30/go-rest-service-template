# Rest Service Boilerplate

Includes:
- Basic configuration [config.go](pkg/config.go)
- Rest service based on [Gin](https://github.com/gin-gonic/gin) and [Fizz](https://github.com/wI2L/fizz) for OpenApi spec generation
- Basic healthcheck (ping handler) (in `common` package)
- Swagger UI available on `http://localhost:8080/swagger` an Open API spec on `http://localhost:8080/openapi.json`
- serving Prometheus metrics on `http://localhost:10250/metrics`
- custom metrics defined in [metrics.go](pkg/metrics/custom_metrics.go)
- unit test for handlers in [pkg/server/router_test.go](pkg/server/router_test.go)
- service-test tests in [/test/stest](test/stest/service_test.go) and infrastructure to run them
- Dockerfiles for service and [service test](Int.Dockerfile)

## How to use 

- Copy whole directory 
- Replace `rest-service` with your service name everywhere in the project. Including Dockerfiles,  Makefile and go.mod
- Adjust paths in go.mod according to the location of your project
- Adjust import paths if needed
- run `go mod tidy`
- run `make build`
- run `make test`

