package main

import (
	"flag"
	"uptime/config"
	"uptime/monitor"

	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()

	var c config.Conf

	confFile := flag.String("conf", "config.yml", "Path to the file that holds the configuration.")
	flag.Parse()
	f := *confFile

	c.GetConfigFile(f)
	
	for _, destination := range c.Destinations {
		baseUrl := destination.BaseUrl
		for _, route := range destination.Routes {
			monitor.SendMonitoringRequest(baseUrl, route)
		}
	}
} 