# ARRAY
Массив - это набор элементов, которые хранятся непрерывным блоком в памяти.

## Особенности:
- элементы хранятся в непрерывном блоке памяти.
- порядок элементов известен, от 0 до n.
- ячейки массива обладают номером(индексом), это логическая структура и память не занимает.
- доступ к ячейке выполняется по индексу.
- у нас есть адрес начала массива.
- у нас есть размер элемента массива.
- чтобы обратиться к элементу id=2, нужно вычислить его адрес и просто обраиться к данному участку памяти.

## Операции:
- Доступ  
Доступ к элементу выполняется очень быстро, за константное время. Т.к. индекс нам известен, остается вычислить адрес => O(1).
- Удаление  
После удаления элемента из массива придется сдвинуть оставшиеся элементы максимум на n шагов => O(n).
- Вставка  
Чтобы вставить элемент в начало, придется сдвинуть остальные элементы, если есть место => O(n).
Когда места в массиве нет, то придется выделить новую память.
Конечно же, вставка в конец осуществляется быстро.
- Поиск  
Для того, чтобы найти элемент в массиве, придется пройтись по всем элементам последовательно => O(n).

## Преимущества
- Компактность. Хранит только данные.
- Быстрый доступ по индексу.
- Не фракментирует память, т.к. данные хранятся одним блоком.

## Недостатки
- Медленные вставка и удаление.