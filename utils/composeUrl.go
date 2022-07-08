package utils

import "strings"

// CreateFullUrl takes a base url and a path and merges them
// together to a full url. It smartly removes any trailing slashes
// so that no invalid url will be returned.
func CreateFullUrl(baseUrl string, slug string) string {
	sanizizedBaseUrl := strings.TrimRight(baseUrl, "/")
	sanitizedSlug := strings.TrimLeft(slug, "/")

	return sanizizedBaseUrl + "/" + sanitizedSlug
}