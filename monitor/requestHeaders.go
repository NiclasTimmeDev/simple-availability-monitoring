package monitor

import (
	"net/http"
	"strings"
	"uptime/config"
)


func GetHeadersFromConfig(route config.Route) []string{
	return route.Headers
}

func AddRequestHeaders(req *http.Request, route config.Route){
	headers := GetHeadersFromConfig(route)


	for _,h := range headers {
		split := strings.Split(h, ":")
		req.Header.Set(split[0], split[1])
	}
}