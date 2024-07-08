# Binary Heap


## Info
Бинарная куча (binary heap) — это структура данных, которая представляет собой полное бинарное дерево, 
в котором каждый узел удовлетворяет определенному свойству кучи. Это свойство может быть либо свойством минимальной кучи (min-heap), 
либо свойством максимальной кучи (max-heap).

Бинарные кучи часто используются в алгоритмах, требующих эффективного доступа к минимальному или максимальному элементу, 
таких как алгоритмы сортировки (например, пирамидальная сортировка), 
алгоритмы на графах (например, алгоритм Дейкстры для поиска кратчайших путей) и других.

Бинарная куча обычно реализуется с использованием массива, что обеспечивает эффективное использование памяти и быстрый доступ к элементам. 
Пространственная сложность самой структуры кучи составляет O(n), где n — количество элементов в куче, так как каждый элемент хранится в массиве.

В минимальной куче (min-heap) каждый узел имеет значение, которое меньше или равно значениям его потомков. 
В максимальной куче (max-heap) каждый узел имеет значение, которое больше или равно значениям его потомков.


## Основные операции
Основные операции с бинарной кучей включают:
- Вставка элемента (insert) — добавление нового элемента в кучу с сохранением свойства кучи.
- Извлечение минимального или максимального элемента (extract-min или extract-max) — удаление корневого элемента кучи и восстановление свойства кучи.
- Проверка, пуста ли куча (is-empty) — определение, содержит ли куча элементы.
- Получение минимального или максимального элемента без удаления (peek) — возвращение корневого элемента кучи без изменения структуры кучи.


## Сложность методов
### Insert
- Временная сложность: O(log n), где n — количество элементов в куче. Это связано с тем, что после вставки элемента в конец кучи, 
может потребоваться "всплытие" (heapify-up) элемента до тех пор, пока не будет восстановлено свойство кучи. 
В худшем случае элемент всплывает до корня, что требует log n шагов.

- Пространственная сложность: O(1), так как вставка элемента не требует дополнительной памяти, 
кроме небольшого количества переменных для хранения индексов.


### Extract-min or Extract-max
- Временная сложность: O(log n). После удаления корневого элемента, последний элемент кучи перемещается в корень, 
и затем может потребоваться "просеивание" (heapify-down) элемента до тех пор, пока не будет восстановлено свойство кучи. 
В худшем случае элемент просеивается до самого нижнего уровня, что требует log n шагов.

- Пространственная сложность: O(1), так как извлечение элемента не требует дополнительной памяти, 
кроме небольшого количества переменных для хранения индексов.


### IsEmpty
- Временная сложность: O(1). Проверка пустоты кучи заключается в проверке количества элементов в куче.
- Пространственная сложность: O(1), так как эта операция не требует дополнительной памяти.

### Peek
- Временная сложность: O(1). Получение корневого элемента кучи не изменяет структуру кучи и не требует дополнительных операций.
- Пространственная сложность: O(1), так как эта операция не требует дополнительной памяти.