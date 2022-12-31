package is_prefix_array

// IsPrefixArray - used for check that string is prefix of array.
func IsPrefixArray(s string, words []string) bool {
	buf := ""
	for _, w := range words {
		buf += w
		if buf == s {
			return true
		}
	}

	return false
}
