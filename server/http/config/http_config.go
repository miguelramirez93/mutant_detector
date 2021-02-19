package config

import "os"

var defaultServerPort = "8080"

// HTTPGetServerPort gets the http server port configuration
func HTTPGetServerPort() string {
	serverPort, ok := os.LookupEnv("PORT")
	if !ok {
		serverPort = defaultServerPort
	}
	return serverPort
}
