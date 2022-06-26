package main

import (
	"interview/app"
	"interview/conf"
)

// GRPC and REST

// Swagger for REST, proto for GRPC

// POST ping endpoint for Postman (should contain object message)
// Fields in return object: echo, timestamp,

func main() {
	configuration := conf.GetConf()

	app.StartService(configuration)
}
