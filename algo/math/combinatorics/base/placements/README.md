# Размещения

## Определение
Пусть имеется n различных объектов.
Будем выбирать из них m объектов и переставлять всеми возможными способами между собой (то есть меняется и состав выбранных объектов, и их порядок). Получившиеся комбинации называются размещениями из n объектов по m, а их число равно: n!/(n-m)! 

## Ньюансы
1. Если количество объектов равно количеству элементов в размещении, т.е. n = m, то результат размещения эквивалентен результату перестановки, т.е. факториалу.


## Example-1
Пример всех размещений из n=3 объектов (различных фигур) по m=2 - на картинке справа. Согласно формуле, их должно быть:
3!/(3-2)! = 1*2*3/1! = 1*2*3/1 = 6/1 = 6. Т.е. из 3х элементов можно составить 6 размещений/пар.