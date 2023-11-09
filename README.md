# Rest Service Boilerplate

Includes:
- Basic configuration [config.go](pkg/config.go)
- Rest service based on [Gin](https://github.com/gin-gonic/gin) and [Fizz](https://github.com/wI2L/fizz) for OpenApi spec generation
- Ping request handler (in `common` package)
- Swagger UI check route `/swagger`
- serving Prometheus metrics on ":10250/metrics". 
- custom metrics defined in [metrics.go](pkg/metrics/metrics.go)
- unit test for handlers in (pkg/server/router_test.go)[pkg/server/router_test.go]
- service-test tests in (/test/stest)(test/stest) and infrastructure to run them
- Dockerfiles for service and [service test](Int.Dockerfile)

## How to use 

- Copy whole directory 
- Replace `rest-service` with your service name everywhere in the project. Including Dockerfiles,  Makefile and go.mod
- Adjust paths in go.mod according to the location of your project
- Adjust import paths if needed
- run `go mod tidy`
- run `make build`
- run `make test`

