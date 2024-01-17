# 783. Minimum Distance Between BST Nodes


## Level - easy


## Task
Given the root of a Binary Search Tree (BST), return the minimum difference between the values of any two different nodes in the tree.


## Объяснение
Задача заключается в нахождении минимального расстояния между узлами бинарного дерева поиска (BST).

Бинарное дерево поиска (BST) - это бинарное дерево, где для всех узлов удовлетворяется свойство: 
все узлы в левом поддереве имеют значения меньше, а все узлы в правом поддереве имеют значения больше, 
чем значение узла-родителя.

Задача требует написать функцию, которая будет находить минимальное расстояние между любыми двумя узлами в BST. 
Расстояние между двумя узлами - это абсолютная разница между их значениями.

## Example 1:
![img.png](img.png)
````
Input: root = [4,2,6,1,3]
Output: 1
````

## Example 2:
![img_1.png](img_1.png)
````
Input: root = [1,0,48,null,null,12,49]
Output: 1
````


## Constraints:
- The number of nodes in the tree is in the range [2, 100].
- 0 <= Node.val <= 10^5
