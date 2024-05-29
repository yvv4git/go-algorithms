# 589. N-ary Tree Preorder Traversal


## Level - easy


## Task
Given the root of an n-ary tree, return the preorder traversal of its nodes' values.

Nary-Tree input serialization is represented in their level order traversal. 
Each group of children is separated by the null value (See examples).


## Объяснение
Задача предполагает, что у вас есть дерево, в котором каждый узел может иметь от 0 до N дочерних узлов. 
Ваша задача - выполнить обход дерева в прямом порядке (preorder traversal), то есть сначала посещать корень, 
затем обходить левое поддерево, а затем правое.

Вам нужно будет реализовать функцию, которая будет принимать корень дерева и возвращать список значений узлов в порядке, 
указанном выше.

Например, если у вас есть дерево:
````
      1
    / | \
   3  2  4
  / \
 5   6
````

Ваша функция должна вернуть [1, 3, 5, 6, 2, 4].

Эта задача часто используется в соответствующих алгоритмах и структурах данных, таких как деревья и графы.


## Example 1:
![img.png](img.png)
````
Input: root = [1,null,3,2,4,null,5,6]
Output: [1,3,5,6,2,4]
````


## Example 2:
![img_1.png](img_1.png)
````
Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
Output: [1,2,3,6,7,11,14,4,8,12,5,9,13,10]
````


## Constraints:

- The number of nodes in the tree is in the range [0, 10^4].
- 0 <= Node.val <= 10^4
- The height of the n-ary tree is less than or equal to 1000.