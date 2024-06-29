# 295. Find Median from Data Stream


## Level - medium


## Task
The median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value, 
and the median is the mean of the two middle values.
- For example, for arr = [2,3,4], the median is 3.
- For example, for arr = [2,3], the median is (2 + 3) / 2 = 2.5.

Implement the MedianFinder class:
- MedianFinder() initializes the MedianFinder object.
- void addNum(int num) adds the integer num from the data stream to the data structure.
- double findMedian() returns the median of all elements so far. Answers within 10-5 of the actual answer will be accepted.


## Объяснение
Задача требует от нас разработать структуру данных, которая поддерживает следующие две операции:
- addNum(int num): Добавляет целое число в структуру данных.
- findMedian(): Возвращает медиану текущего набора данных.

Медиана — это среднее значение в упорядоченном наборе данных. 
Если количество элементов нечетное, медиана — это центральный элемент. 
Если количество элементов четное, медиана — это среднее арифметическое двух центральных элементов.

Для решения этой задачи можно использовать два кучи (heap): одна максимальная куча для хранения меньшей половины элементов 
и одна минимальная куча для хранения большей половины элементов. 
Это позволит нам эффективно поддерживать медиану при добавлении новых элементов.



## Example 1:
````
Input
["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
[[], [1], [2], [], [3], []]
Output
[null, null, null, 1.5, null, 2.0]

Explanation
MedianFinder medianFinder = new MedianFinder();
medianFinder.addNum(1);    // arr = [1]
medianFinder.addNum(2);    // arr = [1, 2]
medianFinder.findMedian(); // return 1.5 (i.e., (1 + 2) / 2)
medianFinder.addNum(3);    // arr[1, 2, 3]
medianFinder.findMedian(); // return 2.0
````


## -105 <= num <= 105
- There will be at least one element in the data structure before calling findMedian.
- At most 5 * 10^4 calls will be made to addNum and findMedian.