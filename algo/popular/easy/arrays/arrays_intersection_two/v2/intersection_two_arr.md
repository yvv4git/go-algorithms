# Intersection two array

## О задаче
Можно решить сортировкой, за более долгое время, но без выделения дополнительной памяти. А можно выделить дополнительную память и решить за линейное время.
Надо посчитать количество появлений элементов первого массива (лучше брать тот, что покороче) — используем для этого словарь. 
Потом пройтись по второму массиву и вычитать из словаря те элементы, которые есть в нем. По ходу добавляем в результат те элементы, у которых частота появлений больше нуля.