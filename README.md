# Rest Service Boilerplate

## :gift: What is inside
- Rest service based on [Gin](https://github.com/gin-gonic/gin) and [Fizz](https://github.com/wI2L/fizz) for OpenApi spec generation
- Basic healthcheck (ping handler) (in `common` package)
- Swagger UI available on `http://localhost:8080/swagger` an Open API spec on `http://localhost:8080/openapi.json`
- serving Prometheus metrics on `http://localhost:10250/metrics`
- custom metrics defined in [metrics.go](pkg/metrics/custom_metrics.go)
- unit test for handlers in [pkg/server/router_test.go](pkg/server/router_test.go)
- service-test tests in [/test/stest](test/stest/service_test.go). Refer to [service tests](https://github.com/num30/api-integration-test#rest-api-integration-test-example) repo for more info
- Dockerfiles for service and [service test](Int.Dockerfile)
- Basic configuration [config.go](pkg/config.go) based on [config](https://github.com/num30/config) package
- CI with [Github Actions](.github/workflows/build.yaml) 
- Github actions for publishing docker image with versioning( tags set to `github-sha` and `branch-[build-number]`). Separate tags used for application and test image

# :computer: Try it
- run unit tests `go test ./... -v`
- run server `go run cmd/server/main.go` 
- run service tests from another terminal `go test ./test/stest -tags servicetest  -v -count=1`
- Send PUT request:
```
        curl -X 'PUT' 'http://localhost:8080/things/first-thing' \
        -d '{ "value": "thing value" }'  
```
- Send GET request 
```
        curl 'http://localhost:8080/things/first-thing'
```
- Check swagger `http://localhost:8080/swagger`
- Check Open API spec `http://localhost:8080/openapi.json`
- Check Prometheus metrics `http://localhost:10250/metrics`

## :whale: Try Docker Without Cloning the code 
- run service  `docker run -p 8080:8080 -p 10250:10250 orsol/go-rest-service-template:latest`
- try steps from 3 to the end from [Try it](#computer-try-it) section


## :bulb: Hints to get started 

- Copy whole directory 
- Replace `rest-service` with your service name everywhere in the project. Including Dockerfiles,  Makefile and go.mod
- Adjust paths in go.mod according to the location of your project
- Adjust import paths if needed
- run `go mod tidy`
- update routes and handlers in [pkg/server/router.go](pkg/server/router.go) defining API objects in [pkg/server/api.go](pkg/server/api.go)
- Define you configuration structure in [pkg/config.go](pkg/config.go)

