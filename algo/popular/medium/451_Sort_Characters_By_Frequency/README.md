# 451. Sort Characters By Frequency


## Level - medium


## Task
Given a string s, sort it in decreasing order based on the frequency of the characters. The frequency of a character is the number of times it appears in the string.

Return the sorted string. If there are multiple answers, return any of them.


## Объяснение
Задача требует отсортировать символы входной строки в порядке убывания частоты их появления. Частота символа определяется количеством его вхождений в строку.

Ключевые аспекты задачи:
1. Чувствительность к регистру:
   - Символы в разных регистрах ('A' vs 'a') считаются разными
   - Это влияет на подсчет частот и итоговый результат

2. Множественность допустимых решений:
   - Если несколько символов имеют одинаковую частоту, их порядок может быть любым
   - Например, для строки "tree" допустимы оба варианта: "eert" и "eetr"

3. Группировка одинаковых символов:
   - Все одинаковые символы должны идти подряд в результирующей строке
   - Варианты с перемешанными символами (например "cacaca") недопустимы

Подробные подходы к решению:

1. Метод подсчета и сортировки:
   - Создаем хэш-таблицу для подсчета частот каждого символа
   - Сортируем символы по убыванию частоты (при равных частотах порядок не важен)
   - Формируем итоговую строку, повторяя каждый символ соответствующее количество раз

2. Метод блочной сортировки (Bucket Sort):
   - Создаем массив "ведер" (buckets), где индекс соответствует частоте символа
   - Распределяем символы по соответствующим ведрам
   - Обходим ведра от самого большого индекса к меньшему
   - Для каждого ведра добавляем его символы в результат

3. Метод кучи (Priority Queue):
   - Строим максимальную кучу (max-heap) на основе частот символов
   - Последовательно извлекаем символы с наибольшей частотой
   - Добавляем извлеченные символы в результирующую строку
   - Повторяем до опустошения кучи

Сложность алгоритмов:
- Подсчет с сортировкой: O(n log n) из-за сортировки
- Bucket Sort: O(n) при оптимальной реализации
- Priority Queue: O(n log k), где k - количество уникальных символов



## Example 1:
```
Input: s = "tree"
Output: "eert"
Explanation: 'e' appears twice while 'r' and 't' both appear once.
So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
```


## Example 2:
```
Input: s = "cccaaa"
Output: "aaaccc"
Explanation: Both 'c' and 'a' appear three times, so both "cccaaa" and "aaaccc" are valid answers.
Note that "cacaca" is incorrect, as the same characters must be together.
```

## Example 3:
```
Input: s = "Aabb"
Output: "bbAa"
Explanation: "bbaA" is also a valid answer, but "Aabb" is incorrect.
Note that 'A' and 'a' are treated as two different characters.
```
 

## Constraints:
- 1 <= s.length <= 5 * 105
- s consists of uppercase and lowercase English letters and digits.