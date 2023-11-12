package _02_happy_number

func isHappyV1(n int) bool {
	/*
		Time complexity: O(1) или O(log(n)^2)
		Space complexity: O(1)
	*/
	for {
		summa := 0
		for n > 0 { // Поскольку внутренний цикл выполняется внутри внешнего цикла, общий время выполнения будет O(log(n)) * O(log(n)) = O(log(n)^2)
			d := n % 10    // log(n) - это количество цифр в числе n. Это происходит потому, что в каждой итерации цикла мы уменьшаем n в 10 раз.
			summa += d * d // Выполняем константное количество операций.
			n /= 10
		}

		if summa == 1 {
			return true
		}

		if summa == 4 {
			return false
		}
		n = summa
	}
}
