# 98. Validate Binary Search Tree


## Level - medium


## Task
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:
- The left subtree of a node contains only nodes with keys less than the node's key.
- The right subtree of a node contains only nodes with keys greater than the node's key.
- Both the left and right subtrees must also be binary search trees.


## Объяснение
Эта задача на проверку, является ли данное дерево бинарного поиска. 
Дерево бинарного поиска - это дерево, в котором для каждого узла выполняется условие: 
все узлы в левом поддереве меньше значения узла, а все узлы в правом поддереве больше значения узла.


## Example 1:
![img.png](img.png)
``
Input: root = [2,1,3]
Output: true
``


## Example 2:
![img_1.png](img_1.png)
``
Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.
``


## Constraints:
- The number of nodes in the tree is in the range [1, 10^4].
- -2^31 <= Node.val <= 2^31 - 1