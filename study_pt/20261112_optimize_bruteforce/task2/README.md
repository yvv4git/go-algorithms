# B. Courier

## Info

Дано N деревень с координатами на плоскости.
Нужно построить маршрут курьера, который стартует и заканчивается в деревне 1 и посещает каждую деревню ровно один раз.
Цель - минимизировать суммарную евклидову длину маршрута, то есть решить задачу коммивояжёра для N <= 3000 точек.
Точный перебор невозможен, поэтому используют жадные построения и эвристические улучшения, например 2-opt; сложность зависит от выбранного подхода.

## Level - hard

## Task

A musical "Zuid-West" was staged in a certain village.
To increase the number of tickets sold, a ticket office was opened in each of the neighboring villages.
The musical has a courier whose duties include visiting all villages every day, starting from the musical's office in village number 1.
During each trip, the courier delivers tickets to sales points and collects the money earned during the previous day.
After that, the courier returns to the office to hand over the revenue.

You are given the coordinates of the villages on a plane.
It can be assumed that the distance between villages is equal to the straight-line distance between the corresponding points on the plane.
It is required to choose a route for the courier that has the smallest possible length.

### Input

You are given input files for this problem.
Each file contains first the number N (N <= 3000) - the number of villages, and then N pairs of numbers specifying the coordinates of the villages.
The coordinates do not exceed 10^9 in absolute value.

### Output

Output the route.
The route must start and end at village number 1, and each of the villages with numbers 2..N must appear in the route description exactly once.

### Scoring

Each of the 10 tests is worth 10 points.
You get 5 points if your route is shorter than the one produced by the greedy algorithm (always move to the nearest unvisited village).

Let `greedy` be the total length of the route obtained by the greedy algorithm.
Let `user` be the length of the route in your solution, and `best` be the length of the route in the best solution.

The remaining points are distributed according to the formula `5 * min(1, (greedy - user) / (0.9 * (greedy - best)))`.

If your solution receives the status FAIL, it means that your solution is better than the one currently stored in the system.
We will fix this.

## Объяснение

Курьер должен объехать все деревни по одному разу и вернуться в офис, который находится в деревне 1.
Маршрут можно записать как последовательность номеров деревень, которая начинается и заканчивается числом 1.
Например, маршрут `1 3 2 4 1` означает: старт в 1, затем 3, затем 2, затем 4, затем возврат в 1.

Расстояние между деревнями равно расстоянию по прямой между точками на плоскости.
Если деревни имеют координаты `(x1, y1)` и `(x2, y2)`, то расстояние между ними равно `sqrt((x1 - x2)^2 + (y1 - y2)^2)`.
Длина маршрута - это сумма длин всех переходов по порядку вдоль маршрута.

Это задача коммивояжёра: нужно найти самый короткий цикл, проходящий через все вершины ровно один раз.
При N <= 3000 точный алгоритм перебора всех перестановок неприменим, поэтому на практике строят маршрут жадно и улучшают его локальными перестановками.

Возможные подходы:

- Greedy nearest neighbor (жадный ближайший сосед). На каждом шаге идти в ближайшую ещё не посещённую деревню. Время O(N^2), память O(N).
- 2-opt improvement (локальное улучшение 2-opt). Переворачивать участки, если маршрут стал короче. Время O(N^2) на итерацию, память O(N).
- Simulated annealing (имитация отжига). Случайно менять маршрут и иногда принимать ухудшения. Время ограничено лимитом запуска, память O(N).

## Example 1

```text
Input:
4
0 0
1 1
0 1
1 0

Output:
1 3 2 4 1
```

Explanation: the courier starts at village 1 with coordinates (0, 0), visits villages 3, 2, 4, and returns to 1.

## Example 2

```text
Input:
10
1 2
3 4
5 6
7 8
9 10
11 12
13 14
15 16
17 18
19 20

Output:
1 2 3 4 5 6 7 8 9 10 1
```

Explanation: all villages lie on one line in increasing order, so visiting them in order 2..10 gives a short route.

## Constraints

- N <= 3000
- -10^9 <= x_i <= 10^9
- -10^9 <= y_i <= 10^9

## См. также

- [Подходы и их сравнение](approachs.md)
- [Greedy nearest neighbor](solve1_greedy_nearest_neighbor)
- [2-opt](solve2_two_opt)
- [Simulated annealing](solve3_simulated_annealing)
