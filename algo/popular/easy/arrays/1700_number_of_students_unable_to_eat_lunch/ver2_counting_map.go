package _700_number_of_students_unable_to_eat_lunch

func countStudentsV2(students []int, sandwiches []int) int {
	/*
		METHOD: Counting map
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)
	*/
	studMap := make(map[int]int)
	for i := 0; i < len(students); i++ {
		studMap[students[i]]++
	}

	unEat := len(students)
	for i := 0; i < len(sandwiches); i++ {
		if studMap[sandwiches[i]] == 0 {
			break
		}
		studMap[sandwiches[i]]--
		unEat--
	}

	return unEat
}
