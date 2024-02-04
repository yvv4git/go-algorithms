package _700_number_of_students_unable_to_eat_lunch

func countStudentsV1(students []int, sandwiches []int) int {
	/*
		METHOD: Counting slice
		Time complexity: O(n)
		Space complexity: O(1)

		Time complexity
		Временная сложность этого алгоритма составляет O(n), где n - количество студентов.
		Это связано с тем, что мы проходимся по каждому студенту и сендвичу ровно один раз.

		Space complexity
		Пространственная сложность также составляет O(1), так как мы используем фиксированное количество памяти для хранения счетчиков.

		В данном алгоритме мы подсчитываем количество студентов, которые предпочитают сендвичи (1) и количество тех, кто не предпочитает (0).
		Затем мы проходимся по очереди сендвичей и, если у нас есть студент, который предпочитает этот вид сендвича,
		то мы уменьшаем счетчик студентов, которые предпочитают этот вид сендвича.
		Если у нас нет студента, который предпочитает этот вид сендвича, то мы останавливаемся и возвращаем количество оставшихся студентов.


		1. Make an array of ints counters of len == 2, this array will count the number of students that are unable to eat.
		2. We have 2 arrays as input students []int and sandwiches []int, both arrays have a series of 0 or 1.
		3. In the first loop we count all the times a 0 or 1 appears in students[]int and keep the count in our array counters []int. If the value is 0 we increment the value of counters[0] by 1 and viceversa.
		4. In our second loop we iterate through all the values in sandwiches []int and decrement the corresponding value from counters []int and if a certain value we try to decrement is already equal to zero we break because there's not another student that will take the current value from sandwiches. If there's not another student that wants that sandwich then the rest of the students will be unable to eat.
		5. We return the sum of the students left from counters
	*/
	var counters [2]int

	// Подсчитываем количество студентов, которые предпочитают сендвичи (1) и количество тех, кто не предпочитает (0)
	for i := 0; i < len(students); i++ {
		counters[students[i]]++
	}

	// Проходимся по очереди сендвичей
	for i := 0; i < len(sandwiches); i++ {
		// Если у нас нет студента, который предпочитает этот вид сендвича, то останавливаемся
		if counters[sandwiches[i]] == 0 {
			break
		}
		// Уменьшаем счетчик студентов, которые предпочитают этот вид сендвича
		counters[sandwiches[i]]--
	}

	// Возвращаем количество оставшихся студентов
	return counters[0] + counters[1]
}
