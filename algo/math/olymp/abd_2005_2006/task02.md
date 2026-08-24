# 2. Five Fishermen

## Info

Задача на пропорциональное рассуждение: определить время для другого
количества рыбаков и судаков при неизменной производительности.
Ключевая идея - скорость поедания на одного рыбака не меняется.
Сложность O(1).

## Level - easy

## Task

Five fishermen ate five zanders in five days. How many days will it take
ten fishermen to eat ten zanders?

## Объяснение

Пять рыбаков съели пять судаков за пять дней. Это значит, что каждый
рыбак съел одного судака за пять дней (производительность: 1 судак /
рыбак / 5 дней).

Если рыбаков станет десять, а судаков - десять, то каждый рыбак всё
равно должен съесть по одному судаку. Поскольку скорость одного рыбака
не изменилась, на это уйдут те же пять дней.

Альтернативно: за пять дней пять рыбаков съедают пять судаков.
Другие пять рыбаков за те же пять дней съедят ещё пять судаков.
Итого десять рыбаков за пять дней съедят десять судаков.

## Example 1

```text
Input:
(no input)

Output:
5 days

Explanation: Each fisherman eats 1 zander in 5 days. 10 fishermen eat
10 zanders in the same 5 days.
```

## Constraints

- Rate per fisherman is constant
- All fishermen eat at the same rate
- No input required; output is the number of days

## См. также

- [ver1_rate.go](ver1_rate.go)
