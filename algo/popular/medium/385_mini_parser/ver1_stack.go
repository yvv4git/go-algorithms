package main

import (
	"fmt"
)

// NestedInteger представляет собой вложенный список, который может содержать целые числа или другие вложенные списки.
type NestedInteger struct {
	isInteger bool             // Указывает, является ли значение целым числом
	value     int              // Хранит целочисленное значение
	list      []*NestedInteger // Хранит список вложенных объектов
}

// SetInteger устанавливает значение NestedInteger на целое число.
func (n *NestedInteger) SetInteger(value int) {
	n.isInteger = true
	n.value = value
	n.list = nil
}

// Add добавляет вложенный объект в список NestedInteger.
func (n *NestedInteger) Add(elem NestedInteger) {
	n.isInteger = false
	n.list = append(n.list, &elem)
}

// GetList возвращает список вложенных объектов, если NestedInteger содержит список.
// Если NestedInteger содержит одно целое число, длина списка будет равна 0.
func (n NestedInteger) GetList() []*NestedInteger {
	return n.list
}

// IsInteger проверяет, является ли NestedInteger целым числом.
func (n NestedInteger) IsInteger() bool {
	return n.isInteger
}

// GetInteger возвращает целое число, которое хранит NestedInteger.
// Результат не определен, если NestedInteger содержит список.
func (n NestedInteger) GetInteger() int {
	return n.value
}

const eof = rune(0)

// Функция deserialize декодирует строку, представляющую вложенную структуру, в NestedInteger.
func deserialize(s string) *NestedInteger {
	/*
		METHOD: Stack
		TIME COMPLEXITY: O(n), где n - длина строки s
		SPACE COMPLEXITY: O(n), используется дополнительная память на хранение stack
	*/

	// Добавляем символ конца файла (EOF) в конец строки, чтобы гарантировать, что последнее число будет обработано.
	s += string(eof)
	// Создаем фиктивный NestedInteger, который будет содержать результат.
	dummy := new(NestedInteger)
	// Инициализируем стек списков, начиная с фиктивного NestedInteger.
	lists := intStack{dummy}

	// Создаем буфер для сбора чисел.
	num := make([]rune, 0, 10)

	// Функция maybeAddNum добавляет новый NestedInteger с целочисленным значением, если буфер num не пуст.
	maybeAddNum := func() {
		if len(num) == 0 {
			return
		}
		newInt := new(NestedInteger)
		newInt.SetInteger(parseInt(num))
		lists.peek().Add(*newInt)
		num = num[:0]
	}

	// Итерируемся по каждому символу в строке.
	for _, ch := range s {
		switch ch {
		case '[':
			// Если символ '[', то создаем новый NestedInteger и добавляем его в стек.
			lists.push(new(NestedInteger))
		case ']':
			// Если символ ']', то добавляем текущий числовой NestedInteger в стек, если он существует.
			maybeAddNum()
			// Извлекаем текущий список из стека.
			list := lists.pop()
			// Добавляем его в родительский список.
			lists.peek().Add(*list)
		case ',', eof:
			// Если символ ',' или EOF, то добавляем текущее числовое значение в стек, если оно существует.
			maybeAddNum()
		default:
			// Если символ не является специальным, то добавляем его в буфер числа.
			num = append(num, ch)
		}
	}
	// Возвращаем первый (и единственный) элемент фиктивного списка, который содержит результат.
	return dummy.GetList()[0]
}

// Функция parseInt преобразует слайс рун, представляющий число, в целое число.
func parseInt(num []rune) int {
	// Инициализируем знак числа как положительный.
	sign := 1
	// Если первый символ в слайсе '-', то число отрицательное, и меняем знак на -1.
	if num[0] == '-' {
		sign = -1
		// Удаляем символ '-' из слайса.
		num = num[1:]
	}
	// Переменная res будет содержать итоговое целочисленное значение.
	var res int
	// Итерируемся по каждому символу в слайсе num.
	for _, ch := range num {
		// Умножаем текущее значение res на 10, чтобы сделать место для следующей цифры.
		res *= 10
		// Добавляем значение текущей цифры к res, вычитая значение руны '0'.
		res += int(ch - '0')
	}
	// Возвращаем итоговое целочисленное значение с учетом знака.
	return sign * res
}

// Тип intStack представляет собой стек для хранения указателей на NestedInteger.
type intStack []*NestedInteger

// Метод peek возвращает последний элемент стека без его удаления.
func (s *intStack) peek() *NestedInteger {
	// Возвращаем последний элемент в стеке.
	return (*s)[len(*s)-1]
}

// Метод push добавляет новый элемент в стек.
func (s *intStack) push(x *NestedInteger) {
	// Добавляем новый элемент в конец стека.
	*s = append(*s, x)
}

// Метод pop удаляет и возвращает последний элемент стека.
func (s *intStack) pop() *NestedInteger {
	// Получаем количество элементов в стеке.
	n := len(*s)
	// Получаем последний элемент стека.
	el := (*s)[n-1]
	// Удаляем последний элемент из стека.
	*s = (*s)[:n-1]
	// Возвращаем удаленный элемент.
	return el
}

func main() {
	s := "324"
	ni := deserialize(s)
	fmt.Println(ni.GetInteger()) // Вывод: 324

	s = "[123,[456,[789]]]"
	ni = deserialize(s)
	// Вывод вложенного списка
	for _, ni := range ni.GetList() {
		if ni.IsInteger() {
			fmt.Println(ni.GetInteger())
		} else {
			for _, innerNI := range ni.GetList() {
				if innerNI.IsInteger() {
					fmt.Println(innerNI.GetInteger())
				} else {
					for _, innermostNI := range innerNI.GetList() {
						fmt.Println(innermostNI.GetInteger())
					}
				}
			}
		}
	}
}
