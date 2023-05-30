package ex01

// Перестановки.
// Пусть есть n объектов.
// Необходимо определить количество возможных сочетаний элементов.
func TransposeCnt(n int) int {
	return factorial(n)
}

// Используем факториал без рекурсии. Динамическое программирование.
func factorial(n int) int {
	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	result := 1
	for i := 1; i <= n; i++ {
		result = result * i
	}

	return result
}

// Размещения элементов, когда элементы в группе не повторяются.
// Количество должно быть равно n! (факториал).
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

func perm(a []rune, f func([]rune), i int) {
	// fmt.Printf("a:%s i:%d \n", string(a), i)
	if i > len(a) {
		// fmt.Printf("i:%d \n", i)
		f(a) // вызываем функцию, которая должна что-то сделать с слайсом данных, например распечатать.
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		// fmt.Printf("i:%d j:%d \n", i, j)
		// fmt.Println("1>", string(a))
		a[i], a[j] = a[j], a[i]
		// fmt.Println("2>", string(a))
		perm(a, f, i+1) // передав слайс в функцию, там будет создан новый слайс, но с таким же указателем на массив.
		a[i], a[j] = a[j], a[i]
		// fmt.Println("3>", string(a))
	}
}
