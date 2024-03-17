# 147. Insertion Sort List


## Level - medium


## Task
Given the head of a singly linked list, sort the list using insertion sort, and return the sorted list's head.

The steps of the insertion sort algorithm:
1. Insertion sort iterates, consuming one input element each repetition and growing a sorted output list.
2. At each iteration, insertion sort removes one element from the input data, finds the location it belongs within the sorted list and inserts it there.
3. It repeats until no input elements remain.

The following is a graphical example of the insertion sort algorithm. 
The partially sorted list (black) initially contains only the first element in the list. 
One element (red) is removed from the input data and inserted in-place into the sorted list with each iteration.

![img.png](img.png)


## Объяснение
Задача предлагает вам реализовать алгоритм сортировки вставками для связного списка.

Сортировка вставками - это простой алгоритм сортировки, который работает так: 
на каждом шаге алгоритма мы берем один элемент из списка и вставляем его в правильное место в уже отсортированной части списка.

## Example 1:
![img_1.png](img_1.png)
````
Input: head = [4,2,1,3]
Output: [1,2,3,4]
````


## Example 2:
![img_2.png](img_2.png)
````
Input: head = [-1,5,3,4,0]
Output: [-1,0,3,4,5]
````


## Constraints:
- The number of nodes in the list is in the range [1, 5000].
- -5000 <= Node.val <= 5000