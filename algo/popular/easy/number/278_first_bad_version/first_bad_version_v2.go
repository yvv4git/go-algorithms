package _78_first_bad_version

func firstBadVersionV2(n int) int {
	var minBad = n
	var maxGood, curr int
	for {
		curr = (minBad + maxGood) / 2
		if minBad-maxGood == 1 {
			return minBad
		}
		if isBadVersion(curr) && curr < minBad {
			minBad = curr
		} else if curr > maxGood {
			maxGood = curr
		}
	}
}
