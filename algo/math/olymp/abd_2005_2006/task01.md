# 1. Three Boxes Puzzle

## Info

Логическая задача на сопоставление содержимого ящиков с неверными надписями.
Ключевая идея: использовать тот факт, что все надписи ложны, чтобы по очереди
исключать варианты. Сложность O(1) - решается прямым логическим выводом.

## Level - easy

## Task

There are three boxes containing groats, vermicelli, and sugar. The first box
is labeled "groats", the second "vermicelli", and the third "groats or sugar".
The contents of each box do not match its label. Determine what is in each box.

## Объяснение

Поскольку все надписи неверны, анализируем их по очереди:

1. На третьем ящике написано "крупа или сахар". Это утверждение ложно,
   значит, в третьем ящике **ни крупа, ни сахар** - только **вермишель**.

2. На первом ящике написано "крупа". Это ложь, значит, там не крупа.
   Вермишель уже в третьем ящике, поэтому в первом - **сахар**.

3. Во втором ящике написано "вермишель". Это ложь, значит, там не вермишель.
   Сахар в первом, вермишель в третьем - остаётся **крупа**.

Итог: 1-й - сахар, 2-й - крупа, 3-й - вермишель.

## Example 1

```text
Input:
(no input)

Output:
Box 1: sugar
Box 2: groats
Box 3: vermicelli

Explanation: All labels are false. Box 3 label "groats or sugar" is false
→ vermicelli. Box 1 label "groats" is false, and vermicelli is taken → sugar.
Box 2 label "vermicelli" is false → groats.
```

## Constraints

- Exactly 3 boxes with distinct contents: groats, vermicelli, sugar
- Each box label is guaranteed to be incorrect
- No input required; output is the determined assignment

## См. также

- [ver1_logic.go](ver1_logic.go)
