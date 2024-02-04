package binary_number_alternating_bits

type void struct{}

var member void

// Массив, в котором содержаться числа, в которых биты чередуются.
var arr = []int{
	1,
	2,
	10,
	42,
	170,
	682,
	2730,
	10922,
	43690,
	174762,
	699050,
	2796202,
	11184810,
	44739242,
	178956970,
	715827882,
	5,
	21,
	85,
	341,
	1365,
	5461,
	21845,
	87381,
	349525,
	1398101,
	5592405,
	22369621,
	89478485,
	357913941,
	1431655765,
}

func hasAlternatingBitsV1(n int) bool {
	/*
		METHOD: Use constants
		TIME COMPLEXITY: O(1)
		Space complexity: O(1)
	*/
	set := make(map[int]void)
	for _, x := range arr {
		set[x] = member
	}
	_, found := set[n]
	return found
}
