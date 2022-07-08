package monitor

import (
	"net/http"
	"strings"
	"uptime/config"
)

// GetHeadersFromConfig returns the headers for a route
// from the route configuration in config.yml
func GetHeadersFromConfig(route config.Route) []string{
	return route.Headers
}

// AddRequestHeaders adds headers for a path configured in
// config.yml to a http.Request object.
func AddRequestHeaders(req *http.Request, route config.Route){
	headers := GetHeadersFromConfig(route)


	for _,h := range headers {
		split := strings.Split(h, ":")
		req.Header.Set(split[0], split[1])
	}
}