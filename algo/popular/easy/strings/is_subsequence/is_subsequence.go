package is_subsequence

// IsSubsequence - used for check is subsequence.
func IsSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	for i := 0; i < len(t); i++ {
		if s[0] == t[i] {
			return IsSubsequence(s[1:], t[i+1:])
		}
	}

	return false
}
