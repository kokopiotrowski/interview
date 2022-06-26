package app

import (
	"interview/conf"
	"interview/rest"
)

func StartService(configuration *conf.Conf) {

	rest.StartServer(configuration.RestPort)

	// // clean up
	// signalChan := make(chan os.Signal, 1)
	// signal.Notify(signalChan, os.Interrupt)
	// defer func() {
	// 	signal.Stop(signalChan)
	// 	cancel()
	// }()
}
