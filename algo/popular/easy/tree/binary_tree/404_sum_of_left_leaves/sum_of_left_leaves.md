# 404. Sum of Left Leaves


## Level = easy


## Task
Given the root of a binary tree, return the sum of all left leaves.

A leaf is a node with no children. A left leaf is a leaf that is the left child of another node.


## Объяснение
Учитывая корень двоичного дерева, верните сумму всех левых листьев.

Лист - это узел без дочерних элементов. 
Левый лист - это лист, который является левым дочерним элементом другого узла.

Надо посчитать сумму всех левых листов.

Задачу можно решить с помощью DFS или BFS алгоритмов.
DFS предпочтительней если дерево сбалансированное. 
BFS может быть предпочтителен, если узел находится на одном уровне с корнем. 
На мой взгляд, это ни наш случай. Вот если бы надо было найти сумму всех листов
на всех уровнях, то это было бы хорошим алгоритмом.

## Example 1:
![img.png](img.png)
````
Input: root = [3,9,20,null,null,15,7]
Output: 24
Explanation: There are two left leaves in the binary tree, with values 9 and 15 respectively.
````

## Example 2:
````
Input: root = [1]
Output: 0
````

## Constraints:
- The number of nodes in the tree is in the range [1, 1000].
- -1000 <= Node.val <= 1000