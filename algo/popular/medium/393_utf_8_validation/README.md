# 393. UTF-8 Validation


## Level - medium


## Task
Given an integer array data representing the data, return whether it is a valid UTF-8 encoding (i.e. it translates to a sequence of valid UTF-8 encoded characters).

A character in UTF8 can be from 1 to 4 bytes long, subjected to the following rules:
- For a 1-byte character, the first bit is a 0, followed by its Unicode code.
- For an n-bytes character, the first n bits are all one's, the n + 1 bit is 0, followed by n - 1 bytes with the most significant 2 bits being 10.

This is how the UTF-8 encoding would work:
````

     Number of Bytes   |        UTF-8 Octet Sequence
                       |              (binary)
--------------------+-----------------------------------------
1          |   0xxxxxxx
2          |   110xxxxx 10xxxxxx
3          |   1110xxxx 10xxxxxx 10xxxxxx
4          |   11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
````

- x denotes a bit in the binary form of a byte that may be either 0 or 1.
- Note: The input is an array of integers. Only the least significant 8 bits of each integer is used to store the data. 
This means each integer represents only 1 byte of data.


## Объяснение
Задача заключается в проверке, является ли данный массив целых чисел допустимой последовательностью байтов в кодировке UTF-8.
UTF-8 — это кодировка символов, которая может представлять все символы Unicode и использует от 1 до 4 байтов для кодирования каждого символа. 
Структура байтов в UTF-8 выглядит следующим образом:
1. Однобайтовые символы (ASCII):
Первый бит равен 0, остальные 7 бит кодируют символ.
Пример: 0xxxxxxx
2. Двухбайтовые символы:
Первый байт начинается с 110, следующие 5 бит кодируют символ.
Второй байт начинается с 10, следующие 6 бит кодируют символ.
Пример: 110xxxxx 10xxxxxx.
3. Трехбайтовые символы:
Первый байт начинается с 1110, следующие 4 бит кодируют символ.
Второй и третий байты начинаются с 10, следующие 6 бит кодируют символ.
Пример: 1110xxxx 10xxxxxx 10xxxxxx.
4. Четырехбайтовые символы:
Первый байт начинается с 11110, следующие 3 бит кодируют символ.
Второй, третий и четвертый байты начинаются с 10, следующие 6 бит кодируют символ.
Пример: 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx

Задача состоит в том, чтобы написать функцию, которая принимает массив целых чисел, представляющих байты, 
и возвращает true, если этот массив является допустимой последовательностью байтов в кодировке UTF-8, 
и false в противном случае.

Примеры:
1. Вход: [197, 130, 1]  
197 в двоичном виде: 11000101 (начало двухбайтового символа)  
130 в двоичном виде: 10000010 (продолжение двухбайтового символа)  
1 в двоичном виде: 00000001 (однобайтовый символ)  

2. Вывод: true
Вход: [235, 140, 4]  
235 в двоичном виде: 11101011 (начало трехбайтового символа)  
140 в двоичном виде: 10001100 (неправильное продолжение)  
Вывод: false  

Для решения этой задачи нужно пройтись по массиву байтов и проверять, соответствуют ли они правилам UTF-8 кодирования.


## Example 1:
````
Input: data = [197,130,1]
Output: true
Explanation: data represents the octet sequence: 11000101 10000010 00000001.
It is a valid utf-8 encoding for a 2-bytes character followed by a 1-byte character.
````


## Example 2:
````
Input: data = [235,140,4]
Output: false
Explanation: data represented the octet sequence: 11101011 10001100 00000100.
The first 3 bits are all one's and the 4th bit is 0 means it is a 3-bytes character.
The next byte is a continuation byte which starts with 10 and that's correct.
But the second continuation byte does not start with 10, so it is invalid.
````


## Constraints:
- 1 <= data.length <= 2 * 10^4
- 0 <= data[i] <= 255