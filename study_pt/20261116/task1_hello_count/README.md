# 1. Hello Counter

## Info

В опенспейсе работают a интровертов и b экстравертов, которые заходят по одному в случайном порядке.
Ключевая идея - посчитать ответ по парам сотрудников, потому что порядок входа не меняет число приветов.
Итоговая сложность - O(1) по времени и O(1) по памяти.

## Level - easy

## Task

An open space has a introverts and b extroverts.
Every day they come to work in a random order, entering the open space one by one.

When an introvert enters the open space, every extrovert already inside says "Hello!" to him.
The introvert answers each of them "Hello!".
When an extrovert enters the open space, he says "Hello!" to every employee already inside.
Every employee answers him "Hello!".

Each spoken "Hello!" counts as one.
So when one employee says "Hello!" and another answers "Hello!", that exchange counts as 2.

Count the total number of times "Hello!" was said during the day.
Note that the answer does not depend on the order in which the employees enter.

### Input

The input contains two integers a and b (0 <= a, b <= 10^9) - the number of introverts and extroverts.

### Output

Print the total number of times "Hello!" was said during the day.
The answer can be quite large, so use 64-bit variables.

## Объяснение

Интроверт - это сотрудник, который сам не здоровается первым, а отвечает на приветствие.
Экстраверт - это сотрудник, который здоровается со всеми, кто уже в опенспейсе.
Когда в опенспейс входит сотрудник, каждый обмен приветствиями даёт ровно два слова "Hello!":
один говорит, второй отвечает.

Порядок входа не важен, потому что каждая подходящая пара сотрудников обменивается приветствиями ровно один раз за день.

Разберём все три вида пар.

- Пара "интроверт и экстраверт": когда второй из них входит, первый уже внутри, и они обмениваются двумя приветствиями.
- Пара "два экстраверта": когда второй экстраверт входит, первый внутри, и они тоже обмениваются двумя приветствиями.
- Пара "два интроверта": интроверты не начинают разговор, таких приветов нет.

Интровертов a, экстравертов b, значит пар "интроверт и экстраверт" ровно a*b, и каждая даёт 2 привета.
Среди экстравертов пар b*(b-1)/2, и каждая даёт 2 привета, то есть b*(b-1) приветов.
Складываем: ответ равен 2*a*b + b*(b-1).

Проверим на примере a = 1, b = 2.
Пар "интроверт и экстраверт" - 2, они дают 2*2 = 4 привета.
Пар "два экстраверта" - 1, она даёт 2 привета.
Всего 4 + 2 = 6, как в первом примере.

Возможные подходы:

- Formula (замкнутая формула). Сразу вычислить 2*a*b + b*(b-1), время O(1), память O(1).

## Example 1

```text
Input:
1 2

Output:
6
```

Explanation: one introvert and two extroverts.
The introvert exchanges two hellos with each extrovert, that is 4.
The two extroverts exchange two hellos with each other, that is 2.
The total is 4 + 2 = 6.

## Example 2

```text
Input:
5 0

Output:
0
```

Explanation: there are no extroverts, and introverts never greet each other first.
So nobody says "Hello!".

## Constraints

- 0 <= a, b <= 10^9
- The answer fits in a 64-bit signed integer.
- The answer does not depend on the order of arrival.

## См. также

- [Подходы и их сравнение](approachs.md)
- [Formula](solve1_formula)
