# 1370. Increasing Decreasing String


## Level - easy


## Task
You are given a string s. Reorder the string using the following algorithm:
1. Remove the smallest character from s and append it to the result.
2. Remove the smallest character from s that is greater than the last appended character, and append it to the result.
3. Repeat step 2 until no more characters can be removed.
4. Remove the largest character from s and append it to the result.
5. Remove the largest character from s that is smaller than the last appended character, and append it to the result.
6. Repeat step 5 until no more characters can be removed.
7. Repeat steps 1 through 6 until all characters from s have been removed.

If the smallest or largest character appears more than once, you may choose any occurrence to append to the result.

Return the resulting string after reordering s using this algorithm.


## Объяснение
Дана строка s, состоящая из строчных английских букв. Необходимо преобразовать её в новую строку, следуя определённому алгоритму:
- Шаг 1: Выберите наименьший символ из строки s и добавьте его в результат.
- Шаг 2: Выберите следующий наименьший символ из оставшихся символов строки s, который больше последнего добавленного символа, и добавьте его в результат.
- Шаг 3: Повторяйте шаг 2, пока больше нельзя будет выбрать символ.
- Шаг 4: Выберите наибольший символ из оставшихся символов строки s и добавьте его в результат.
- Шаг 5: Выберите следующий наибольший символ из оставшихся символов строки s, который меньше последнего добавленного символа, и добавьте его в результат.
- Шаг 6: Повторяйте шаг 5, пока больше нельзя будет выбрать символ.
- Шаг 7: Повторяйте шаги 1–6, пока все символы из строки s не будут использованы.

Пусть дана строка s = "aaaabbbbcccc".
- Сначала выбираем наименьшие символы в порядке возрастания: a, a, a, a.
- Затем выбираем наибольшие символы в порядке убывания: c, c, c, c.
- Потом снова выбираем наименьшие символы: b, b, b, b.
- В итоге получаем строку "aaaacccbbbb".

Для решения задачи можно использовать следующие шаги:
- Подсчитайте количество каждого символа в строке с помощью массива или хэш-таблицы.
- Постройте результат, чередуя выбор наименьших и наибольших символов, пока все символы не будут использованы.



## Example 1:
```
Input: s = "aaaabbbbcccc"
Output: "abccbaabccba"
Explanation: After steps 1, 2 and 3 of the first iteration, result = "abc"
After steps 4, 5 and 6 of the first iteration, result = "abccba"
First iteration is done. Now s = "aabbcc" and we go back to step 1
After steps 1, 2 and 3 of the second iteration, result = "abccbaabc"
After steps 4, 5 and 6 of the second iteration, result = "abccbaabccba"
```

## Example 2:
```
Input: s = "rat"
Output: "art"
Explanation: The word "rat" becomes "art" after re-ordering it with the mentioned algorithm.
```

## Constraints:
- 1 <= s.length <= 500
- s consists of only lowercase English letters.