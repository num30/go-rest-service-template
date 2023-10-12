package rest

import (
	cors "github.com/gin-contrib/cors"
)

type HttpConfig struct {
	Port int         `default:"8080" envvar:"SERVICE_PORT"`
	Cors cors.Config `default:"{\"AllowOrigins\": [\"http://localhost\", \"https://localhost\", \"*\"]}"`
}
