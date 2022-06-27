package main

import (
	"interview/app"
	"interview/conf"
)


func main() {
	configuration := conf.GetConf()

	app.StartService(configuration)
}
