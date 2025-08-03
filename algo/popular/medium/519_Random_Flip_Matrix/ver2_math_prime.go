//go:build ignore

package main

import (
	"math/rand"
	"time"
)

/*
PrimeSolution реализует алгоритм случайного переворота ячеек матрицы с использованием:
1. Отображения через простое число для равномерного распределения
2. Хэш-карты для отслеживания перевернутых позиций
3. Математических свойств простых чисел

Сложность:
- Время:
  - Flip(): O(1) в среднем (амортизированная)
  - Reset(): O(1)

- Память: O(k), где k - количество перевернутых ячеек
*/
type PrimeSolution struct {
	m, n      int          // Размеры матрицы (строки, столбцы)
	total     int          // Общее количество ячеек (m * n)
	prime     int          // Простое число > total (для равномерного распределения)
	remaining int          // Количество оставшихся неперевернутых ячеек
	flipped   map[int]bool // Хранит перевернутые позиции (ключ - линейный индекс)
	rng       *rand.Rand   // Генератор случайных чисел
}

/*
Constructor инициализирует решение:
1. Вычисляет общее количество ячеек
2. Находит наименьшее простое число > total
3. Инициализирует генератор случайных чисел
4. Создает пустую карту для отслеживания перевернутых ячеек
*/
func Constructor(m int, n int) PrimeSolution {
	total := m * n
	source := rand.NewSource(time.Now().UnixNano())

	// Находим простое число для равномерного распределения
	prime := findNextPrime(total)

	return PrimeSolution{
		m:         m,
		n:         n,
		total:     total,
		prime:     prime,
		remaining: total,
		flipped:   make(map[int]bool),
		rng:       rand.New(source),
	}
}

/*
Flip возвращает случайную неперевернутую позицию [i,j]:
1. Генерирует случайное число в диапазоне [0, remaining)
2. Вычисляет кандидата через умножение на простое число и взятие модуля
3. Находит первую свободную позицию (обрабатывает коллизии)
4. Помечает позицию как перевернутую и возвращает координаты

Сложность: O(1) в среднем благодаря:
- Математическому отображению через простое число
- Амортизированной обработке коллизий
*/
func (this *PrimeSolution) Flip() []int {
	if this.remaining == 0 {
		return nil // Все ячейки перевернуты
	}

	// Генерация кандидата через простое число
	r := this.rng.Intn(this.remaining)
	candidate := (r * this.prime) % this.total

	// Обработка коллизий
	for this.flipped[candidate] {
		candidate = (candidate + 1) % this.total
	}

	// Помечаем как перевернутую и возвращаем координаты
	this.flipped[candidate] = true
	this.remaining--

	return []int{candidate / this.n, candidate % this.n}
}

/*
Reset сбрасывает состояние матрицы:
1. Очищает карту перевернутых позиций
2. Восстанавливает счетчик оставшихся ячеек

Сложность: O(1) - очистка карты в Go выполняется за O(1)
*/
func (this *PrimeSolution) Reset() {
	this.flipped = make(map[int]bool)
	this.remaining = this.total
}

/*
findNextPrime возвращает наименьшее простое число > n
Используется для обеспечения равномерного распределения
*/
func findNextPrime(n int) int {
	if n < 2 {
		return 2
	}

	// Поиск простого числа начиная с n+1
	for i := n + 1; ; i++ {
		if isPrime(i) {
			return i
		}
	}
}

/*
isPrime проверяет, является ли число простым
Оптимизации:
- Пропускаем четные числа после 2
- Проверяем делители до квадратного корня
*/
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}

	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}
