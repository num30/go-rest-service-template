package version

// Minor version is autoupdated by the build system
// NOTE: use go build -ldflags "-X github.com/num30/go-rest-service-template/cmd/version.Version==$(git describe)"
var Version = "v0.0.debug"
