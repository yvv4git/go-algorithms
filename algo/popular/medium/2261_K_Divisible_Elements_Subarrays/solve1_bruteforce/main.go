package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const _subarraySeparator byte = ','

func countDistinct(nums []int, k int, p int) int {
	/*
		TASK: Посчитать количество различных подмассивов массива nums, в которых не более k элементов делятся на p.
		Подмассив - это непрерывная часть массива, заданная границами left и right.
		Два подмассива различны, если у них разная длина или хотя бы один элемент на одинаковой позиции отличается.
		Элемент x делится на p, если x % p == 0.
		Пример: nums = [2, 3, 3, 2, 2], k = 2, p = 2.
		Элементы 2, 2 и 2 делятся на 2, поэтому подмассив [2, 3, 3, 2] подходит, а [2, 3, 3, 2, 2] нет: в нём три таких элемента.
		Различных подходящих подмассивов здесь 11.
		Граничные случаи:
		k = 0: подходит только подмассив без элементов, делящихся на p, например [3] или [3, 3];
		k = n: подходит любой подмассив, и ответ равен числу различных подмассивов.
		Отрицательный пример: nums = [2, 2], k = 1, p = 2, подмассив [2, 2] содержит два делящихся элемента, поэтому он не учитывается, ответ равен 1.
		Функция возвращает количество различных подходящих подмассивов.

		METHOD: Bruteforce (перебор с хеш-таблицей)
		Хеш-таблица - это структура map, которая хранит ключи и проверяет их наличие за O(1) в среднем.
		Перебираем левую границу left от 0 до n-1.
		Для каждой left двигаем правую границу right вправо и поддерживаем счётчик divisible.
		Если nums[right] % p == 0, увеличиваем divisible.
		Когда divisible > k, прекращаем перебор для текущей left: при дальнейшем увеличении right счётчик только растёт.
		Каждый подходящий подмассив кодируется в строку key без неоднозначности: числа записываются через strconv.AppendInt и разделяются запятой.
		Ключ добавляется в map distinct, поэтому одинаковые подмассивы считаются один раз.
		Ответ - размер map distinct.

		TIME COMPLEXITY: O(n^3) в худшем случае.
		Всего O(n^2) подмассивов, а кодирование подмассива длины L занимает O(L), отсюда суммарно O(n^3).

		SPACE COMPLEXITY: O(n^3) в худшем случае.
		В map хранятся ключи всех различных подмассивов, суммарная длина ключей может достигать O(n^3).
	*/
	distinct := make(map[string]struct{})

	for left := range nums {
		var key []byte
		divisible := 0

		for right := left; right < len(nums); right++ {
			if nums[right]%p == 0 {
				divisible++
			}
			if divisible > k {
				break
			}

			key = strconv.AppendInt(key, int64(nums[right]), 10)
			key = append(key, _subarraySeparator)
			distinct[string(key)] = struct{}{}
		}
	}

	return len(distinct)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k, p int
	if _, err := fmt.Fscan(reader, &n, &k, &p); err != nil {
		fmt.Fprintln(os.Stderr, err)

		return
	}

	nums := make([]int, n)
	for index := range nums {
		if _, err := fmt.Fscan(reader, &nums[index]); err != nil {
			fmt.Fprintln(os.Stderr, err)

			return
		}
	}

	fmt.Fprintln(writer, countDistinct(nums, k, p))
}
