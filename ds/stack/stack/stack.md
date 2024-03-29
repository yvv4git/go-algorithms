# STACK
Stack - абстрактный тип данных, организованный по принцыпу LIFO. Это не массивы, это очередь. Примером может быть стопка книг.
Всегда кладем на верх и всегда забираем сверху.

## Особенности
- список элементов организованных по принцыпу lifo(Last In First Out)
- операции: push(element), pop, isEmpty, getTop.

## Сложность:
- добавление элемента на вершину стека(push): O(1)
- удаление элемента с вершины стека(pop): O(1)
- получить количество элементов в стеке(count): O(1)


## Стек вызовов
Когда происходит вызов функции, обработчик программы
переходит к другой области памяти. И чтобы после выхода из
текущей функции знать куда вернуться в стеке сохраняется:
- адрес, который получит управление после выхода из функции
- аргументы функции


## Рекурсия
Рекурсия - это представление некоторого объекта с помощью таких же объектов (как правило меньшего размера).

## Рекурсивная функция
Это функция, которая вызывает себя же, но как правило с другими параметрами.
Если f(n) > 0, то вызываем f(n-1).
Иначе выходим.

## Переполнение стека
Это ситуация, когда слишком много рекурсивных вызовов, они складываются в стек, память, выеделенная на стек
заканчивается. В результате в какой-то момент получим ошибку Stack Overflow.


## База рекурсии
База рекурсии - ситуация, в которой поведение функции тривиально и не требует рекурсивных вызовов.
Т.е. ситуация, когда рекурсивный вызов завершается. Когда программа грамотно спроектирована, то
база рекурсии должна срабатывать корректно. Чтобы не уйти в бесконечную рекурсию.