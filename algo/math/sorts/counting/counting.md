# Сортировка методом "Counting Sort"

## Что это такое?
Сортировка методом **Counting Sort** — это не сравнительный алгоритм сортировки, который работает путем подсчета количества вхождений каждого уникального элемента в массив. Алгоритм основывается на знании диапазона значений элементов массива и создает вспомогательную структуру (массив подсчета), чтобы эффективно сортировать данные.

Процесс работы алгоритма:
1. Находим максимальное значение в массиве для определения размера массива подсчета.
2. Создаем массив подсчета, где каждый индекс соответствует количеству вхождений числа в исходном массиве.
3. Из этого массива подсчета восстанавливаем отсортированный массив.

## Временная сложность
- **Лучший, средний и худший случай**: **O(n + k)**, где `n` — количество элементов в массиве, а `k` — диапазон значений (максимальное значение в массиве). Алгоритм проходит по массиву один раз для подсчета вхождений и второй раз для восстановления отсортированного массива.

## Пространственная сложность
- Пространственная сложность: **O(n + k)**, где `n` — количество элементов в исходном массиве, а `k` — диапазон значений. Это связано с необходимостью хранения массива подсчета.

## Плюсы и минусы

### Плюсы:
- **Высокая эффективность для чисел с ограниченным диапазоном**: когда значения чисел в массиве ограничены, сортировка может работать очень быстро.
- **Линейное время при маленьком диапазоне значений**: если `k` (диапазон значений) значительно меньше `n` (количество элементов), то алгоритм работает очень эффективно.
- **Отсутствие сравнений**: поскольку алгоритм не использует операции сравнения, он может быть быстрее, чем другие методы сортировки для определенных данных.

### Минусы:
- **Неэффективен для больших диапазонов значений**: если диапазон значений слишком велик, использование памяти для массива подсчета может стать проблемой.
- **Требует дополнительных затрат по памяти**: необходимо выделить память для массива подсчета, что может быть непрактично для больших диапазонов значений.
- **Только для целых чисел или категориальных данных**: сортировка работает только для целых чисел или данных, которые могут быть представлены в виде счетных величин.

## Где используется?
- **Сортировка чисел с ограниченным диапазоном**: особенно эффективно, когда элементы массива имеют ограниченный диапазон значений (например, возраст людей или количество товаров в магазине).
- **Генерация распределений**: используется для построения частотных распределений или гистограмм.
- **Алгоритмы для целых чисел**: применяется в задачах, где необходимо отсортировать только целые числа.

