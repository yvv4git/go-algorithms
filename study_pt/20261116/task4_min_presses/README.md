# D. Physical Buttons

## Info

Телефон имеет M кнопок, на которых по порядку нужно разместить N эмодзи, причём один эмодзи выбирается числом нажатий на свою кнопку.
Ключевая идея - разбить последовательность эмодзи на M частей подряд так, чтобы сумма весов по позициям внутри кнопок была минимальной.
Итоговая сложность - O(N^2*M) по времени и O(N*M) по памяти.

## Level - medium

## Task

When developing a new phone, engineers came up with the idea of adding physical buttons for entering popular emojis.
Engineers calculated how many times each emoji is used while using the phone and fixed the order in which the emojis must be placed.

There may be more emojis than buttons on the phone.
Therefore, several first emojis will have to be placed on the first button, several next emojis on the second button, and so on.
To select an emoji placed on a button, you need to press it one or more times.
If the emojis "skull", "clown", and "nail polish" are placed on a button, then the first press selects "skull".
The second press selects "clown", and the third press selects "nail polish".

You need to distribute the emojis among the buttons so that the total number of button presses for typing all emojis is minimal.

### Input

The first line contains two integers N and M (1 <= N <= 30, 1 <= M <= 10) - the number of emojis and the number of buttons on the phone keyboard.
The next line contains N integers a_i (0 <= a_i <= 1000000) - the number of times the i-th emoji is used.

### Output

Print M non-negative integers b_i - the number of emojis assigned to the i-th button of the phone.
The sum of all b_i must be equal to N.

If there are several possible answers, print any of them.

## Объяснение

Эмодзи идут в фиксированном порядке и их нельзя переставлять.
На каждой кнопке лежит несколько эмодзи подряд: первый выбирается одним нажатием, второй - двумя, третий - тремя и так далее.
Если на кнопку положили k эмодзи с частотами a_1, a_2, ..., a_k по порядку, то эта кнопка даёт вклад a_1 + 2*a_2 + ... + k*a_k нажатий.
Нужно разбить все N эмодзи на M групп подряд (некоторые группы могут быть пустыми) так, чтобы сумма вкладов всех кнопок была минимальной.

Частые эмодзи выгодно ставить ближе к началу группы, потому что множитель нажатий растёт.
Но порядок задан заранее, поэтому остаётся только выбрать границы групп.
Например, для N = 5, M = 3 и частот 1, 3, 5, 1, 1 можно взять группы (1), (3), (5, 1, 1).
Тогда сумма равна 1*1 + 3*1 + 5*1 + 1*2 + 1*3 = 14.

Возможные подходы:

- Bruteforce (полный перебор разбиений). Перебрать все способы разбить N эмодзи на M групп подряд и выбрать минимум.
  Время O(C(N+M-1,M-1)*N), память O(N).
- Dynamic programming (динамическое программирование). Считать минимальную стоимость для первых i эмодзи на j кнопках и восстановить ответ.
  Время O(N^2*M), память O(N*M).

При N <= 30 и M <= 10 оба подхода укладываются в ограничения, но динамическое программирование проще масштабируется.

## Example 1

```text
Input:
5 3
1 3 5 1 1

Output:
1 1 3
```

Explanation: the first button gets the first emoji, the second gets the second emoji, and the third gets the remaining three.
The total number of presses is 1*1 + 3*1 + 5*1 + 1*2 + 1*3 = 14.

## Constraints

- 1 <= N <= 30
- 1 <= M <= 10
- 0 <= a_i <= 1000000
- The emojis must stay in the given order.
- The sum of all b_i must be equal to N.

## См. также

- [Подходы и их сравнение](approachs.md)
- [Bruteforce](solve1_bruteforce)
- [Dynamic programming](solve2_dp)
