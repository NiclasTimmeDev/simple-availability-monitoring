package httpClient

import (
	"net/http"
	"time"
)

// NewHttpClient creates and configures an http.Client object.
func NewHttpClient() http.Client {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	return client
}