# 703. Kth Largest Element in a Stream


## Level - easy


## Task
Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Implement KthLargest class:
- KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of integers nums.
- int add(int val) Appends the integer val to the stream and returns the element representing the kth largest element in the stream.


## Объяснение
Это задача, в которой требуется найти k-й наибольший элемент в потоке данных.

Решение этой задачи может быть реализовано с помощью кучи (heap) в Go. 
Куча - это структура данных, которая позволяет быстро получать наибольший элемент из набора данных. 
В данном случае, мы будем использовать минимальную кучу, чтобы хранить k наибольших элементов.


## Что такое куча
Куча (Heap) - это один из видов структур данных, который представляет собой дерево, 
где каждый родительский узел имеет значение не меньше (или не больше), чем значения его потомков. 
Это может быть минимальное (или максимальное) куча, в которой родительский узел всегда меньше (или больше) своих потомков.

Кучи часто используются в алгоритмах, которые нуждаются в быстром доступе к наименьшему (или наибольшему) элементу, 
а также в реализации приоритетной очереди.

В Go, кучи реализуются с помощью пакета container/heap, который предоставляет функции для создания, 
вставки, извлечения и удаления элементов из кучи.


## Что такое Min Heap

Min Heap (Минимальная куча) - это структура данных, где родительский узел всегда меньше или равен своим потомкам. 
Это значит, что наименьший элемент всегда находится в корне (или вершине) кучи.

Max Heap (Максимальная куча) - это структура данных, где родительский узел всегда больше или равен своим потомкам. 
Это значит, что наибольший элемент всегда находится в корне (или вершине) кучи.

В обоих случаях, если мы хотим найти наименьший (или наибольший) элемент, мы можем это сделать за константное время, так как он всегда находится в корне.
Min Heap используется, когда нам нужно быстро найти наименьший элемент, а Max Heap - когда нам нужно быстро найти наибольший элемент.


## Example 1:
````
Input
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
Output
[null, 4, 5, 5, 8, 8]

Explanation
KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
kthLargest.add(3);   // return 4
kthLargest.add(5);   // return 5
kthLargest.add(10);  // return 5
kthLargest.add(9);   // return 8
kthLargest.add(4);   // return 8
````

## Constraints:
- 1 <= k <= 10^4
- 0 <= nums.length <= 10^4
- -10^4 <= nums[i] <= 10^4
- -10^4 <= val <= 10^4
- At most 10^4 calls will be made to add.
- It is guaranteed that there will be at least k elements in the array when you search for the kth element.