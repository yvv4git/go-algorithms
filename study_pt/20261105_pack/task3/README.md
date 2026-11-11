# C. Packing and Unpacking

## Info

Задача с двойным запуском: программа один раз сжимает строку из маленьких английских букв, а второй раз распаковывает её обратно.
Ключевая идея - адаптивное арифметическое кодирование с контекстом из одной или двух предыдущих букв, которое даёт меньше 4 бит на букву.
Итоговая сложность - O(n) по времени и O(n) по памяти.

## Level - hard

## Task

### Statement

This is a problem with two runs.
Your solution will be run twice.

You need to implement a compression and decompression algorithm for English texts.
The texts contain only lowercase English letters.
They have no spaces, punctuation marks, or any other characters except `a..z`.
These texts are works of fiction (natural text).

The solution is considered correct if packing followed by unpacking recovers the original string.
The number of bytes in the compressed sequence must also be at most half the number of characters in the original string.

### Interaction Protocol

First run - packing (`pack`):

```text
stdin:
    pack
    <string s of lowercase English letters>

stdout:
    N
    b0 b1 b2 ... b(N-1)
```

- the first line contains N - the number of bytes in the compressed sequence;
- the second line contains N integers from 0 to 255 separated by spaces - the bytes themselves;
- at the end you must flush the output buffer and print a newline.

Second run - unpacking (`unpack`):

```text
stdin:
    unpack
    N
    b0 b1 b2 ... b(N-1)

stdout:
    <original string s>
```

Important: the second run receives exactly the output that the program printed during the first run.
The numbers are not a hint and not a different format, but simply the bytes compressed by your program.
The jury only adds the word `unpack` on top.

Scheme of one test:

```text
   FIRST RUN                    BETWEEN RUNS                 SECOND RUN
+------------------+          +------------------+          +------------------+
| input: pack      |          | output of the    |          | input: unpack    |
|       s          |  ---->   | first run is     |  ---->   |       N          |
| output: N, bytes |          | taken: N numbers |          |       N numbers  |
+------------------+          +------------------+          | output: s        |
                                                             +------------------+
```

### Length

The length of the original string in all tests is from 100000 to 200000 letters.
The number of compressed bytes must be at most half the string length.
In other words, the condition `N <= len(s) / 2` must hold.

Note: this is the same constraint as "no more than half".
The wording "at most half" and the condition `N <= len(s) / 2` mean the same thing.

## Объяснение

### Что вообще происходит

Программа не может запомнить строку между запусками, потому что это два независимых процесса.
Значит, вся информация, нужная для восстановления, должна лежать в тех N байтах, которые программа напечатала в первый раз.
Это и есть сжатие без потерь: из N байт обратно получается ровно исходная строка s.

### Почему нельзя просто вывести строку как есть

Каждая буква в строке - это один символ.
Если печатать байты как есть, понадобится примерно по одному байту на букву.
А лимит - половина байта на букву.
Значит, в среднем каждая буква должна занимать меньше 4 бит.

### Почему обычный Хаффман может не подойти

Если считать частоты всех букв в тексте и кодировать каждую букву отдельно, получится около 4.1-4.2 бита на букву.
Это больше 4 бит, то есть больше половины байта.
В лимит `N <= n/2` такой код не всегда влезает.

Причина в том, что буквы в английском тексте не независимы.
Следующая буква сильно зависит от предыдущих.
Например, после `q` почти всегда идёт `u`, а после `th` очень часто идёт `e`.
Если использовать этот факт, неопределённость следующей буквы падает, и на неё нужно меньше бит.
Именно поэтому в решении используется контекст - одна или две предыдущие буквы.

### Что нужно от алгоритма

1. Быть сжатием без потерь: `unpack(pack(s)) == s` для любой строки из `a..z`.
2. Давать в среднем меньше 4 бит на букву на естественном английском тексте.
3. Быть самодостаточным: всё, что нужно декодеру, лежит внутри N байт.
4. Уметь работать с длиной до 200000 символов за разумное время.

Конкретные алгоритмы и их сравнение разобраны в файле [Подходы и их сравнение](approachs.md).

### Что лежит внутри сжатых байт

В любом решении первые несколько байт обычно отводят под служебную информацию.
Как минимум нужно передать длину исходной строки n, чтобы декодер знал, когда остановиться.
Например, можно записать n четырьмя байтами (little-endian), а дальше уже идут сами сжатые данные.

```text
+------------------+-------------------------------------------+
| длина n (4 байта)| сжатые данные (сколько получилось)        |
+------------------+-------------------------------------------+
```

Длину нельзя не передавать: в сжатом потоке нет явного символа "конец строки".
Без n непонятно, сколько букв восстанавливать.

### Что важно не перепутать

- Два запуска - это два разных процесса, поэтому нельзя хранить состояние в памяти между ними.
- В первый запуск на входе слово `pack`, во второй - `unpack`, и регистр важен.
- Второй запуск получает ровно ваш вывод: N, затем N чисел, и ничего додумывать не нужно.
- Сжатие должно быть без потерь: восстановленная строка должна совпасть символ в символ.
- Лимит считается в байтах, а не в символах: `N <= len(s) / 2`.
- Числа печатаются в десятичной системе, через пробел, от 0 до 255.

## Example 1

Пример из условия показывает только формат обмена, а не конкретный алгоритм сжатия.
Числа `11 255 0 13 253` - это просто произвольные байты, которые как будто выдала программа.

Первый запуск:

```text
Input:
pack
abacabaca

Output:
5
11 255 0 13 253
```

Второй запуск - на вход подаётся вывод первого запуска с добавленным словом `unpack`:

```text
Input:
unpack
5
11 255 0 13 253

Output:
abacabaca
```

Смысл примера: что бы вы ни напечатали в первый раз, во второй раз вы обязаны из этих же чисел восстановить исходную строку.

## Constraints

- 100000 <= len(s) <= 200000.
- Строка состоит только из символов `a`..`z`.
- N - целое число, каждый байт от 0 до 255.
- N <= len(s) / 2.
- Вывод второго запуска должен совпадать с исходной строкой символ в символ.
- Нужно сбрасывать буфер вывода и печатать перевод строки в конце.

## См. также

- [Подходы и их сравнение](approachs.md)
- [Арифметическое кодирование, контекст 1](solve1_arithmetic_order1)
- [Арифметическое кодирование, контекст 2](solve2_arithmetic_order2)
