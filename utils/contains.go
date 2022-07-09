package utils

// StringSliceContains checks if a specific string (needle) is present
// in a slice of strings (haystack). Returns true if it is, else false.
func StringSliceContains(haystack []string, needle string ) bool{
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}