package utils

import "strings"

func CreateFullUrl(baseUrl string, slug string) string {
	sanizizedBaseUrl := strings.TrimRight(baseUrl, "/")
	sanitizedSlug := strings.TrimLeft(slug, "/")

	return sanizizedBaseUrl + "/" + sanitizedSlug
}