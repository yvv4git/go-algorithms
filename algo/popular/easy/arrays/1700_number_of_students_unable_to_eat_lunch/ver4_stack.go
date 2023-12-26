package _700_number_of_students_unable_to_eat_lunch

func countStudentsV4(students []int, sandwiches []int) int {
	for len(students) > 0 && len(sandwiches) > 0 {
		if students[0] == sandwiches[0] {
			students = students[1:]
			sandwiches = sandwiches[1:]
		} else {
			students = append(students[1:], students[0])
		}
	}

	return len(students)
}
