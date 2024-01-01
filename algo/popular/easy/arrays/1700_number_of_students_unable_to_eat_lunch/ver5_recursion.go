package _700_number_of_students_unable_to_eat_lunch

// Функция countStudentsV5 решает задачу о том, сколько студентов останется в очереди, если каждый студент будет пытаться взять сендвич из стека.
// Функция принимает два параметра: students - очередь студентов, sandwiches - стек сендвичей.
// Функция возвращает количество студентов, которые не смогли взять сендвич из-за того, что все остальные студенты хотели быть первыми в очереди.
func countStudentsV5(students []int, sandwiches []int) int {
	/*
		Method: Recursion
		Time complexity: O(n)
		Space complexity: O(n)

		Time complexity
		Временная сложность этого алгоритма составляет O(n), где n - количество студентов.
		Это связано с тем, что мы проходим по всем студентам только один раз.

		Space complexity
		Пространственная сложность этого алгоритма составляет O(n), так как мы храним всех студентов в памяти.
	*/

	// Переменная wantTopSandwich отвечает на вопрос: "Хочет ли первый студент сендвич, который находится в верхней части стека?"
	wantTopSandwich := false
	// Проходим по всем студентам
	for _, v := range students {
		// Если студент хочет сендвич, который находится в верхней части стека, то меняем значение переменной wantTopSandwich на true
		if v == sandwiches[0] {
			wantTopSandwich = true
		}
	}
	// Если ни один студент не хочет сендвич, который находится в верхней части стека, то возвращаем количество студентов
	if !wantTopSandwich {
		return len(students)
	}

	// Если первый студент хочет сендвич, который находится в верхней части стека, то он берет его и сендвич удаляется из стека
	if students[0] == sandwiches[0] {
		return countStudentsV5(students[1:], sandwiches[1:])
	} else {
		// Если первый студент не хочет сендвич, который находится в верхней части стека, то он идет в конец очереди
		return countStudentsV5(append(students[1:], students[0]), sandwiches)
	}
}