package first_unique_character_in_string

func firstUniqCharV3(s string) int {
	/*
		METHOD: Bruteforce.
		Time complexity : O(n^2)
		Space complexity : O(1)
	*/
loop:
	for i := range s {
		for j := range s {
			if i != j && s[i] == s[j] {
				continue loop
			}
		}
		return i
	}
	return -1
}
