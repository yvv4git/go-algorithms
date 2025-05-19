package main

func circularArrayLoop(nums []int) bool {
	/*
		APPROACH: Two Pointers
		1. Используем два указателя (slow и fast) для обхода массива.
		2. На каждом шаге перемещаем slow на один элемент вперед (slow = (slow + nums[slow] + n) % n),
		а fast на два элемента вперед (fast = (fast + nums[fast] + n) % n).
		3. Если встречаем цикл, возвращаем true.
		4. Если находимся в том же направлении и fast догоняет slow, возвращаем false.
		5. Если находимся в разных направлениях, обновляем направление движения.
		6. Повторяем шаги 2-5, пока не достигнем конца массива.
		7. Если не найдем цикл, возвращаем false.

		TIME COMPLEXITY: O(n)
		- В худшем случае мы проходим по всем элементам массива.

		SPACE COMPLEXITY: O(1)
		- Используем то же количество переменных, независимо от размера массива.
	*/
	n := len(nums)
	if n == 0 {
		return false
	}

	for i := 0; i < n; i++ {
		// Если текущий элемент уже посещен (помечен как 0), пропускаем
		if nums[i] == 0 {
			continue
		}

		// Определяем направление движения: вперед (true) или назад (false)
		direction := nums[i] > 0
		slow, fast := i, i

		for {
			// Функция для получения следующего индекса с учетом кругового массива
			next := func(index int) int {
				step := nums[index]
				// Вычисляем следующий индекс с учетом перехода через границы массива
				nextIndex := (index + step) % n
				if nextIndex < 0 {
					nextIndex += n
				}
				return nextIndex
			}

			// Двигаем медленный указатель на один шаг
			slow = next(slow)
			// Проверяем, не изменилось ли направление или не попали в петлю из одного элемента
			if (nums[slow] > 0) != direction || nums[slow] == 0 || next(slow) == slow {
				break
			}

			// Двигаем быстрый указатель на два шага
			fast = next(fast)
			if (nums[fast] > 0) != direction || nums[fast] == 0 || next(fast) == fast {
				break
			}
			fast = next(fast)
			if (nums[fast] > 0) != direction || nums[fast] == 0 || next(fast) == fast {
				break
			}

			// Если медленный и быстрый указатели встретились, значит есть петля
			if slow == fast {
				return true
			}
		}

		// Помечаем все посещенные элементы как 0, чтобы не проверять их снова
		slow = i
		for {
			nextStep := nums[slow]
			nums[slow] = 0 // Помечаем как посещенный
			nextSlow := (slow + nextStep) % n
			if nextSlow < 0 {
				nextSlow += n
			}
			// Останавливаемся, если следующий элемент уже посещен или направление меняется
			if nums[nextSlow] == 0 || (nums[nextSlow] > 0) != direction {
				break
			}
			slow = nextSlow
		}
	}

	return false
}
