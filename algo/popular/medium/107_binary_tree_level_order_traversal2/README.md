# 107. Binary Tree Level Order Traversal II


## Level - medium


## Task
Given the root of a binary tree, return the bottom-up level order traversal of its nodes' values. 
(i.e., from left to right, level by level from leaf to root).


## Объяснение
Задача заключается в том, чтобы обойти двоичное дерево уровней в порядке "снизу вверх". 
Это значит, что вы должны пройтись по каждому уровню дерева, начиная с нижнего, и вернуть результат в виде списка списков, 
где каждый вложенный список содержит значения узлов на одном уровне.

Для решения этой задачи, вы можете использовать алгоритм обхода в ширину (BFS), но вместо того, 
чтобы добавлять узлы на каждом уровне в результирующий список, вы добавляете их в очередь. 
Затем, когда вы закончите обход всех узлов на одном уровне, вы можете добавить их в результирующий список.


## Example 1:
![img.png](img.png)
````
Input: root = [3,9,20,null,null,15,7]
Output: [[15,7],[9,20],[3]]
````


## Example 2:
````
Input: root = [1]
Output: [[1]]
````

## Example 3:
````
Input: root = []
Output: []
````

## Constraints:
- The number of nodes in the tree is in the range [0, 2000].
- -1000 <= Node.val <= 1000