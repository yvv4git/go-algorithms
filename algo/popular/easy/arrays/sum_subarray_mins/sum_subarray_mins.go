package sum_subarray_mins

func sumSubarrayMins(arr []int) int {
	/*
	   Time Complexity: O(N), where N is the length of A.
	   Space Complexity: O(N).
	*/
	out := 0
	var st stack
	st.push(-1) // setting left bound

	for i := range arr {
		for len(st) > 1 && arr[i] <= arr[st.peak()] {
			idx := st.pop()
			out += (i - idx) * (idx - st.peak()) * arr[idx]
		}
		st = append(st, i)
	}

	for len(st) > 1 {
		idx := st.pop()
		out += (len(arr) - idx) * (idx - st.peak()) * arr[idx]
	}

	return out % 1000000007 // 10^9 + 7
}

type stack []int

func (s *stack) push(v int) {
	*s = append(*s, v)
}

func (s *stack) pop() int {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return v
}

func (s *stack) empty() bool {
	return len(*s) == 0
}

func (s *stack) peak() int {
	v := (*s)[len(*s)-1]

	return v
}
