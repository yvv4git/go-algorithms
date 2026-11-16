# 3. Deadlines

## Info

Каждый день Вася решает ровно одну задачу, у каждой задачи есть дедлайн, а за каждый день просрочки начисляется штраф.
Ключевая идея - решать задачи в порядке возрастания дедлайна (правило EDD) и назначать дни 1..N по этому порядку.
Итоговая сложность - O(N log N) по времени и O(N) по памяти.

## Level - medium

## Task

Vasya's job is to solve algorithmic problems.
He solves one algorithmic problem per day.

In the coming days he has to solve N algorithmic problems.
Each problem has a deadline - the number of the day by which this problem must be solved (inclusive).
For each day a problem is overdue past its deadline, Vasya is fined one ruble.

Determine the order in which Vasya should solve the problems so that the fine is minimal.

### Input

The first line contains the number of problems N (1 <= N <= 100).
The second line contains N integers separated by spaces - the deadlines of the problems (1 <= t_i <= 1000000).

### Output

Print N numbers - the numbers of the days on which Vasya should solve the corresponding problem.
The day numbers must not exceed N.
If there are several correct solutions, print any of them.

## Объяснение

Дедлайн задачи - это номер дня, до которого включительно её нужно решить.
Если задачу с дедлайном t решить в день d, то штраф за неё равен max(0, d - t).
Нужно назначить каждой задаче свой день от 1 до N так, чтобы суммарный штраф был минимальным.

Ключевое наблюдение: задачи выгодно решать в порядке возрастания дедлайна.
Это правило называют EDD (Earliest Due Date, самый ранний дедлайн первым).

Почему это работает, покажем обменом на маленьком примере.
Пусть две задачи с дедлайнами 5 и 2 стоят на днях 2 и 3.
Если сначала решить задачу с дедлайном 5, штрафы равны max(0, 2-5) = 0 и max(0, 3-2) = 1, всего 1.
Если сначала решить задачу с дедлайном 2, штрафы равны max(0, 2-2) = 0 и max(0, 3-5) = 0, всего 0.
Видно, что задача с меньшим дедлайном должна стоять раньше.
Повторяя такой обмен, любое расписание можно привести к порядку возрастания дедлайна, не увеличив штраф.

Отсюда алгоритм:

- Сохраняем для каждой задачи пару (дедлайн, номер в порядке ввода).
- Сортируем пары по возрастанию дедлайна, при равенстве порядок любой.
- Задаче на позиции k, считая с 1, назначаем день k.
- Записываем назначенные дни в ответ по исходным номерам и выводим ответ.

Проверим на первом примере: дедлайны равны 2, 3, 1.
После сортировки порядок такой: задача 3 (дедлайн 1), задача 1 (дедлайн 2), задача 2 (дедлайн 3).
Назначаем дни 1, 2, 3: задача 3 получает день 1, задача 1 - день 2, задача 2 - день 3.
В порядке ввода дни равны 2, 3, 1, штраф равен 0.

Возможные подходы:

- Naive selection (наивный выбор минимума). Каждый день брать задачу с минимальным дедлайном, время O(N^2), память O(N).
- Greedy EDD (сортировка по дедлайну). Отсортировать задачи по возрастанию дедлайна и выдать дни 1..N, время O(N log N), память O(N).

## Example 1

```text
Input:
3
2 3 1

Output:
2 3 1
```

Explanation: after sorting by deadline the order is problem 3, problem 1, problem 2.
They get days 1, 2, 3, so in input order the days are 2, 3, 1.
Every problem is solved on time, so the fine is 0.

## Example 2

```text
Input:
4
2 1 2 1

Output:
4 1 3 2
```

Explanation: this is one optimal assignment.
Problems 2 and 4 have deadline 1 and get days 1 and 2.
Problems 1 and 3 have deadline 2 and get days 3 and 4.
The fines are 2, 0, 1, 1, so the total fine is 4.

## Constraints

- 1 <= N <= 100
- 1 <= t_i <= 1000000
- Vasya solves exactly one problem per day.
- The answer days are numbered from 1 to N.

## См. также

- [Подходы и их сравнение](approachs.md)
- [Naive selection](solve1_naive_selection)
- [Greedy EDD](solve2_greedy_edd)
