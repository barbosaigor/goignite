package config

import (
	"log"

	"github.com/jpfaria/goignite/pkg/config"
)

const Port = "http.server.port"
const StatusRoute = "http.server.route.status"
const HealthRoute = "http.server.route.health"

func init() {

	log.Println("getting configurations for http server")

	config.Add(Port, 8080, "server http port")
	config.Add(StatusRoute, "/resource-status", "define status url")
	config.Add(HealthRoute, "/health", "define health url")

}

func GetPort() int {
	return config.Int(Port)
}

func GetStatusRoute() string {
	return config.String(StatusRoute)
}

func GetHealthRoute() string {
	return config.String(HealthRoute)
}