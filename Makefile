.PHONY: service-test.test test build lint deps 

build: stest.build
	@echo Building version: \"$(BIN_VER)\"
	env CGO_ENABLED=0 go build -ldflags='-X github.com/num30/go-rest-service-template/cmd/version.Version=$(BIN_VER) -extldflags=-static' -o bin/rest-service cmd/server/main.go

test:
	go test -v ./...

# Go lint
lint:
	golangci-lint run

deps:
	go install

clean:
	rm pb/gen/*.go


# service tests
stest.test:
	go test ./test/stest -tags stest  -v -count=1


stest.build:
	env CGO_ENABLED=0 go test ./test/stest -tags stest -v -c -o bin/service-test