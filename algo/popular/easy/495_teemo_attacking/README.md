# 495. Teemo Attacking


## Level - easy


## Task
Our hero Teemo is attacking an enemy Ashe with poison attacks! When Teemo attacks Ashe, Ashe gets poisoned for a exactly duration seconds. 
More formally, an attack at second t will mean Ashe is poisoned during the inclusive time interval [t, t + duration - 1]. 
If Teemo attacks again before the poison effect ends, the timer for it is reset, and the poison effect will end duration seconds after the new attack.

You are given a non-decreasing integer array timeSeries, where timeSeries[i] denotes that Teemo attacks Ashe at second timeSeries[i], and an integer duration.

Return the total number of seconds that Ashe is poisoned.


## Объяснение
В задаче описывается ситуация, где персонаж Тимо (Teemo) атакует врага с определенной периодичностью. Вам даны два массива:
- timeSeries: Массив целых чисел, где каждый элемент представляет время, в которое Тимо наносит удар по врагу.
- duration: Целое число, представляющее продолжительность действия яда от одного удара.

Задача состоит в том, чтобы определить общее время, в течение которого враг будет под действием яда.


## Example 1:
```
Input: timeSeries = [1,4], duration = 2
Output: 4
Explanation: Teemo's attacks on Ashe go as follows:
- At second 1, Teemo attacks, and Ashe is poisoned for seconds 1 and 2.
- At second 4, Teemo attacks, and Ashe is poisoned for seconds 4 and 5.
Ashe is poisoned for seconds 1, 2, 4, and 5, which is 4 seconds in total.
```

Объяснение:
```
Вход: timeSeries = [1, 4], duration = 2
Выход: 4
Объяснение: 
- В момент времени 1 Тимо наносит удар, яд действует до момента 3 (1 + 2).
- В момент времени 4 Тимо наносит еще один удар, яд действует до момента 6 (4 + 2).
- Общее время действия яда: (3 - 1) + (6 - 4) = 2 + 2 = 4.
```


## Example 2:
```
Input: timeSeries = [1,2], duration = 2
Output: 3
Explanation: Teemo's attacks on Ashe go as follows:
- At second 1, Teemo attacks, and Ashe is poisoned for seconds 1 and 2.
- At second 2 however, Teemo attacks again and resets the poison timer. Ashe is poisoned for seconds 2 and 3.
Ashe is poisoned for seconds 1, 2, and 3, which is 3 seconds in total.
```

Объяснение:
```
Вход: timeSeries = [1, 2], duration = 2
Выход: 3
Объяснение: 
- В момент времени 1 Тимо наносит удар, яд действует до момента 3 (1 + 2).
- В момент времени 2 Тимо наносит еще один удар, яд действует до момента 4 (2 + 2).
- Однако, так как второй удар наносится до того, как первый яд перестанет действовать, второй яд просто продлевает действие первого яда до момента 4.
- Общее время действия яда: (3 - 1) + (4 - 3) = 2 + 1 = 3.
```


## Constraints:
- 1 <= timeSeries.length <= 104
- 0 <= timeSeries[i], duration <= 107
- timeSeries is sorted in non-decreasing order.