# 844. Backspace String Compare


## Level - easy


## Task
Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.

Note that after backspacing an empty text, the text will continue empty.


## Объяснение
Задача требует от нас сравнить две строки после того, как все символы # в обеих строках будут обработаны как символы "backspace". 
Символ "backspace" удаляет предыдущий символ в строке, если он существует.

Пример-1
```
Ввод: s = "ab#c", t = "ad#c"
Вывод: true
Объяснение: Обе строки становятся "ac".
```

Пример-2.
```
Ввод: s = "ab##", t = "c#d#"
Вывод: true
Объяснение: Обе строки становятся пустыми: "".
```

Пример-3.
```
Ввод: s = "a##c", t = "#a#c"
Вывод: true
Объяснение: Обе строки становятся "c".
```

Пример-4.
```
Ввод: s = "a#c", t = "b"
Вывод: false
Объяснение: Строка s становится "c", а строка t остается "b".
```


Подход-1.
Мы можем использовать стек для обработки каждой строки. 
Мы будем проходить по каждой строке и добавлять символы в стек, если они не являются #. 
Если встречаем #, мы удаляем верхний элемент стека, если он не пуст.

Подход-2.  
Мы можем использовать два указателя, чтобы пройтись по строкам с конца к началу. 
Мы будем пропускать символы, которые будут удалены backspace операциями, и сравнивать оставшиеся


## Example 1:
```
Input: s = "ab#c", t = "ad#c"
Output: true
Explanation: Both s and t become "ac".
```


## Example 2:
```
Input: s = "ab##", t = "c#d#"
Output: true
Explanation: Both s and t become "".
```


## Example 3:
```
Input: s = "a#c", t = "b"
Output: false
Explanation: s becomes "c" while t becomes "b".
```


## Constraints:
- 1 <= s.length, t.length <= 200
- s and t only contain lowercase letters and '#' characters.


## Follow up: 
Can you solve it in O(n) time and O(1) space?