# 6. Zigzag Conversion


## Level - medium


## Task
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: 
(you may want to display this pattern in a fixed font for better legibility)

````
P   A   H   N
A P L S I I G
Y   I   R
````

And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:
string convert(string s, int numRows);


## Что за ZigZag формат?
Алгоритм Zigzag используется для преобразования строки в "zigzag" формат, 
который часто используется в задачах, связанных с паттернами в строках.

1. Чтение строк
"Zigzag" формат упрощает чтение строк, так как символы в строках "zigzag" перемещаются вверх и вниз,
что делает ее более интуитивно понятной для человека.

2. Оптимизация памяти. 
Алгоритм Zigzag может быть полезен, если вам нужно сохранить строку в "zigzag" формате, чтобы сэкономить память.
Строки с большим количеством пробелов весят много, а в zigzag не хранятся пробелы. Тоже касается
повторяющихся символов. Еще когда среди символов мало уникальных символов.

3. Алгоритмы машинного обучения. 
Алгоритмы машинного обучения, такие как нейронные сети, часто используют "Zigzag" формат для обработки изображений.

4. Кодирование и декодирование.
Алгоритм Zigzag может использоваться для кодирования и декодирования строк, 
что может быть полезно в задачах, связанных с сетевым программированием.

5. Алгоритмы сжатия данных. 
Алгоритмы сжатия данных, такие как Huffman coding, 
могут использовать "Zigzag" формат для более эффективного сжатия данных.

6. Графические интерфейсы. 
В графических интерфейсах "Zigzag" формат может использоваться для создания интересных и необычных графических элементов.

7. Игры. 
В некоторых игре "Zigzag" формат может использоваться для создания интересных и необычных уровней.


## Example 1:
````
Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
````

## Example 2:
````
Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I
````

## Example 3:
````
Input: s = "A", numRows = 1
Output: "A"
````

## Constraints:
- 1 <= s.length <= 1000
- s consists of English letters (lower-case and upper-case), ',' and '.'.
- 1 <= numRows <= 1000