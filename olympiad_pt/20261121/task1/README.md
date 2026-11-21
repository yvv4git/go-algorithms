# A. Сортировка с обманщиком

## Info

Интерактивная задача: нужно восстановить скрытую перестановку чисел от $1$ до $N$, задавая вопросы сравнения двух элементов.
Жюри может ответить неверно не более $k$ раз, причём ложные ответы выбираются адаптивно по ходу взаимодействия.
Ключевая идея - сравнивать элементы с повторными запросами и использовать устойчивую к ошибкам сортировку.

## Level - hard

## Task

The jury has chosen a permutation $\pi = (\pi(1), \pi(2), \ldots, \pi(N))$ of numbers from $1$ to $N$.
The permutation is fixed in advance and does not change during the interaction.
Your task is to reconstruct it by asking comparison queries.

In one query you specify two different indices $x$ and $y$ ($1 \le x, y \le N$).
The jury answers whether $\pi(x) < \pi(y)$ is true.
However, the jury may lie: an answer is called false if it does not match the actual relation between $\pi(x)$ and $\pi(y)$.
During the whole interaction the jury will give at most $k$ false answers.

The jury decides which queries to answer falsely during the interaction and may take your previous queries into account.
In different tests the jury acts differently: in some tests false answers are chosen randomly, in others they are chosen adversarially.
There are no other restrictions on the choice of false answers.
In particular, it is not guaranteed that in every test the permutation can be reconstructed for sure within the query limit.

You may make at most $Q$ queries.
The same query (including with swapped indices) can be asked several times; answers to identical queries may differ.
After that you must output the permutation $\pi$.
The answer is considered correct if it exactly matches the hidden permutation.

### Interaction Protocol

At the beginning of the interaction your program must read a line containing three integers $N$, $k$, and $Q$.
Here $2 \le N \le 1000$, $0 \le k \le 30$, $1 \le Q \le 11000$.

To make a query, output a line `? x y` ($1 \le x, y \le N$, $x \neq y$).
In response the jury outputs `YES` if (according to the jury) $\pi(x) < \pi(y)$, and `NO` otherwise.

When you have reconstructed the permutation, output a line `! p1 p2 ... pN`, where $p_i = \pi(i)$.
The numbers $p_1, \ldots, p_N$ must form a permutation of numbers from $1$ to $N$.
The answer is output exactly once and is not counted as a query; after that your program must terminate.
The answer is accepted if the output permutation matches the hidden one.

If a query is invalid (wrong format, $x = y$, index out of range) or the limit of $Q$ queries is exceeded, the jury outputs `-1`.
In this case your program must terminate immediately - otherwise you will get Wrong Answer.
If the program continues running, the verdict may be any.

After printing each line, flush the output buffer:

- `fflush(stdout)` or `cout.flush()` in C++;
- `System.out.flush()` in Java;
- `sys.stdout.flush()` in Python;
- `flush(output)` in Pascal.

Otherwise you may get Idleness Limit Exceeded.

### Scoring

The sample test is worth $0$ points.
Besides it, there are $50$ tests, each scored independently.

If on a test your program outputs the correct permutation using $t \le Q$ queries, you get

$$1.5 + 0.5 \cdot \frac{Q - t}{Q}$$

points for this test.
If the answer is wrong or the query limit is violated, you get $0$ points for the test.

Points for tests are summed.
Tests differ in values of $N$, $k$, and $Q$, and in the strategy by which the jury chooses false answers.

## Объяснение

Нам неизвестна перестановка $\pi$, то есть каждое значение от $1$ до $N$ стоит на какой-то позиции ровно один раз.
Запрос `? x y` спрашивает, верно ли, что значение на позиции $x$ меньше значения на позиции $y$.
Жюри может соврать не более $k$ раз за всё взаимодействие, а выбирать, когда соврать, может адаптивно.

Главная сложность - обычное сравнение ненадёжно.
Если один раз получили ответ `YES`, он может быть ложным.
Поэтому один и тот же запрос имеет смысл повторять: если ответы противоречат друг другу, значит среди них есть ложь.

Рассмотрим пример из условия: $N = 3$, $k = 1$, $Q = 11000$, скрытая перестановка $\pi = (2, 3, 1)$.
Первые три запроса дали `YES`, `YES`, `YES` для пар $(1,2)$, $(3,1)$, $(2,3)$.
Третий ответ ложный, потому что на самом деле $\pi(2) = 3 > 1 = \pi(3)$.
Но по первым трём ответам перестановку восстановить нельзя: подходит и $(2,3,1)$ с одной ложью, и $(1,2,3)$ с одной ложью.
Повторные запросы `? 3 1` и `? 1 2` дают правдивые `YES`, и тогда вариант $(1,2,3)$ потребовал бы уже двух ложных ответов.
Так как $k = 1$, остаётся только $(2,3,1)$.

Возможные подходы:

- Repeated comparisons (повторные сравнения). Сравнивать пару элементов несколько раз и принимать решение большинством голосов.
- Fault-tolerant sorting (устойчивая к ошибкам сортировка). Использовать сортировку, где каждое сравнение делается через надёжную процедуру.
- Adaptive verification (адаптивная проверка). Построить гипотезу перестановки, затем перепроверять подозрительные сравнения запросами.

## Example 1

```text
Input:
3 1 11000

Output:
? 1 2

Input:
YES

Output:
? 3 1

Input:
YES

Output:
? 2 3

Input:
YES

Output:
? 3 1

Input:
YES

Output:
? 1 2

Input:
YES

Output:
! 2 3 1
```

Explanation: the hidden permutation is $\pi = (2, 3, 1)$.
The third answer is false because actually $\pi(2) = 3 > 1 = \pi(3)$, but the jury answered `YES`.
All other answers are truthful, and $k = 1$ means there can be no other false answers.
The repeated queries resolve the ambiguity and leave only $(2, 3, 1)$.

## Constraints

- $2 \le N \le 1000$
- $0 \le k \le 30$
- $1 \le Q \le 11000$
- The permutation is fixed in advance and does not change during the interaction.
- The jury gives at most $k$ false answers during the whole interaction.
- The same query can be asked several times, and answers to identical queries may differ.

## См. также

- [Подходы и их сравнение](approachs.md)
