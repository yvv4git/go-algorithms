# A. Correcting One Error

## Info

Задача с двойным запуском: при первом запуске программа получает x и выдаёт y с избыточностью, а при втором по z восстанавливает x.
Ключевая идея - добавить к данным контрольные биты по коду Хэмминга, который находит и исправляет одиночную ошибку.
Итоговая сложность - O(n) по времени и O(n) по памяти.

## Level - hard

## Task

This is a problem with two runs.
Your solution will be run twice.

During the first run it receives a binary string x of length at most n.
The program must output a binary string y of length at most m.
The value m is not known to your program.
It depends on the subtask, and the corresponding value is given in the scoring table.
If your program outputs a string longer than m, it will get a Wrong answer verdict.

Between the runs the jury program makes at most one modification in y.
It replaces a zero with a one or a one with a zero and obtains a string z.

During the second run the program receives z and must restore the original string x and print it.

### Input

During the first run the first line contains the number 1.
The second line contains a binary string x of length n (10 <= n <= 100000).

During the second run the first line contains the number 2.
The second line contains a binary string z of length at most m (10 <= m <= 100017).
It is guaranteed that z equals the string y printed during the first run.
It may also be obtained from y by changing exactly one character to the opposite one.

### Output

During the first run print a string y that allows restoring x after a modification is made to it.
The length of y must not exceed m, and m does not exceed n by more than ceil(log2 n).

During the second run, using the given string z, restore the original string x and print it.

## Объяснение

Задача состоит из двух запусков одной и той же программы, между которыми жюри может испортить один бит.
При первом запуске мы получаем исходную строку x и должны выдать строку y, в которой кроме данных есть избыточные контрольные биты.
При втором запуске мы получаем строку z, отличающуюся от y не более чем в одной позиции, и обязаны восстановить x.
То есть мы строим помехоустойчивый код, исправляющий одну ошибку: любая одиночная замена бита не должна мешать декодированию.
Число m ограничивает длину y, поэтому добавлять слишком много избыточности нельзя.

Возможные подходы:

- Repetition code (троирование битов): каждый бит пишется трижды, ошибка исправляется большинством.
- Repetition code имеет время O(n), но длина 3n превышает m.
- Hamming code (код Хэмминга): добавляем k контрольных битов с условием 2^k >= n + k + 1, и синдром укажет номер ошибки.
- Hamming code работает за время O(n).

Код Хэмминга работает так.
Биты строки y нумеруются с 1.
Позиции, номера которых являются степенями двойки (1, 2, 4, 8, ...), отводятся под контрольные биты, а остальные позиции - под данные x.
Контрольный бит с номером 2^j отвечает за чётность всех позиций, у которых j-й бит номера равен единице.
При втором запуске для каждой группы считается чётность, и из полученных значений собирается синдром - номер испорченной позиции.
Если синдром равен нулю, ошибки нет.
Иначе бит в позиции с этим номером инвертируется, после чего из y выписываются информационные биты.
Так как ошибка ровно одна, синдром всегда указывает на неё однозначно.

## Example 1

```text
Input:
1
01010

Output:
000111000111000
```

Explanation: here each bit of x is repeated three times, so x = "01010" becomes y = "000111000111000".

## Example 2

```text
Input:
2
000101000111000

Output:
01010
```

Explanation: the jury changed one bit of y, and majority voting in every triple restores the original x = "01010".

## Constraints

- 10 <= n <= 100000
- 10 <= m <= 100017
- The length of y must not exceed m.
- m does not exceed n by more than ceil(log2 n).
- z equals y or is obtained from y by changing exactly one character to the opposite one.

## См. также

- [Подходы и их сравнение](approachs.md)
- [Repetition code](solve1_repetition)
- [Hamming code](solve2_hamming)
