# E. Transaction Ordering

## Info

Дан набор транзакций, каждая читает и пишет в локации памяти, и нужно выбрать как можно более длинную последовательность без конфликтов.
Ключевая идея - жадно добавлять транзакции в порядок, если их чтения не пересекаются с уже записанными локациями.
Итоговая сложность - O(S) по времени и O(m + n) по памяти, где S - суммарный размер всех множеств.

## Level - hard

## Task

You are developing a module for selecting transactions included in a block in a new blockchain.

The module receives a set of n transactions numbered from 1 to n.
Transactions are executed in a special virtual machine.
The memory of this machine consists of m locations numbered from 1 to m.
For each location, you know in advance the set of memory locations from which this transaction reads and the set of memory locations to which it writes.
For the i-th transaction, denote these sets as RS_i and WS_i.

From this set of transactions, the module must choose a subset and order it.
Let the module choose k transactions with numbers p_1, p_2, ..., p_k.
For speculative execution, there must be no conflicts in the selected sequence of transactions.
A conflict is a pair of transactions where the transaction that comes earlier writes to some memory location that another transaction later reads.
More formally, transactions p_i and p_j where i < j conflict if there exists a memory location l such that l is in WS_{p_i} and l is in RS_{p_j}.

Given this set of transactions, you need to choose a sequence of transactions of the largest possible length without conflicts.
You do not need to find the optimal solution in this problem.
Your solution will receive points proportional to the length of the sequence you find.

### Input

The first line of the input contains two integers n and m - the number of transactions and memory locations respectively (1 <= n <= 10^5, 1 <= m <= 10^6).
Next, descriptions of transactions follow in increasing order of numbers.

Each transaction description contains three lines.
The first of them contains two integers r_i and w_i - the sizes of sets RS_i and WS_i respectively (1 <= r_i, w_i <= m).
The second line contains r_i distinct integers - the elements of the set RS_i.
The third line contains w_i distinct integers - the elements of the set WS_i.

The total size of all sets RS and WS does not exceed 10^6.

### Output

In the first line, print the number k - the number of transactions in the selected sequence.
In the second line, print k distinct integers - the numbers of transactions in the sequence.

### Scoring

For each test in this problem, you can get a real number of points from 0 to 4.

If the sequence you output contains conflicting transactions, you will get 0 points for the test.
Otherwise, you will get 4k/sqrt(n) points.

There are 25 tests in the problem in total.
Thus, it is impossible to get 100 points for this problem.

## Объяснение

Транзакция - это операция, которая читает значения из некоторых ячеек памяти и записывает значения в другие ячейки.
Ячейка памяти - это локация с номером от 1 до m.
Множество чтений RS_i содержит номера локаций, которые i-я транзакция читает.
Множество записей WS_i содержит номера локаций, в которые i-я транзакция пишет.

Последовательность транзакций безопасна, если ни одна более ранняя транзакция не записывает в локацию, которую более поздняя транзакция читает.
Другими словами, для пары на позициях i < j конфликт есть только в случае "запись раньше, чтение позже" по одной и той же локации.
Чтение раньше записи, чтение раньше чтения и запись раньше записи конфликтом не считаются.

Нужно вывести длинную последовательность без конфликтов.
Оптимальное решение не требуется, поэтому можно использовать простые жадные стратегии, которые всегда дают корректный ответ.

Например, в примере tx1 читает 1 и пишет 2, tx2 читает 2 и пишет 3, tx3 читает 3 и пишет 1.
Последовательность (1, 3) безопасна: tx1 пишет только в 2, а tx3 читает 3, пересечения нет.
Последовательность (1, 2) была бы конфликтом, потому что tx1 пишет в 2, а tx2 потом читает 2.

Возможные подходы:

- Empty sequence (пустая последовательность). Всегда выводить 0 транзакций, ответ корректен, но даёт 0 баллов.
- Greedy append (жадное добавление). Проходить транзакции в порядке номеров и добавлять безопасную транзакцию.
  Транзакция безопасна, если её чтения не пересекаются с уже записанными локациями.
  Время O(S), память O(m + n).
- Greedy sorted append (жадное добавление с сортировкой). Отсортировать транзакции по эвристике и применять то же правило.
  Время O(S + n log n), память O(m + n).

## Example 1

```text
Input:
3 3
1 1
1
2
1 1
2
3
1 1
3
1

Output:
2
1 3
```

Explanation: transactions 1 and 3 do not conflict in this order, because transaction 1 writes to location 2 and transaction 3 reads from location 3.
The sequence 1, 2 would be a conflict, because transaction 1 writes to location 2 and transaction 2 later reads from location 2.

## Constraints

- 1 <= n <= 10^5
- 1 <= m <= 10^6
- 1 <= r_i, w_i <= m
- All elements inside each RS_i and each WS_i are distinct.
- The total size of all sets RS and WS does not exceed 10^6.
- Any conflict-free sequence is a valid answer for correctness.

## См. также

- [Подходы и их сравнение](approachs.md)
- [Empty sequence](solve1_empty)
- [Greedy append](solve2_greedy)
- [Greedy sorted append](solve3_greedy_sorted)
