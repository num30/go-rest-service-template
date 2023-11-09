name: test-and-lint
on:
  push:

jobs:
  test:
    name: test
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v4

    - name: Set up Go
      uses: actions/setup-go@v2
      with:
        go-version: 1.21

    - name: Test
      run:  go build -v ./...

    - name: Test
      run:  go test -v ./... -coverprofile="coverage.out"

  lint:
    name: lint
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v4

    - name: Set up Go
      uses: actions/setup-go@v2
      with:
        go-version: 1.21

    - name: golangci-lint
      uses: golangci/golangci-lint-action@v3.1.0     