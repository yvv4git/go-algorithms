# 1089. Duplicate Zeros


## Level - easy


## Task
Given a fixed-length integer array arr, duplicate each occurrence of zero, shifting the remaining elements to the right.

Note that elements beyond the length of the original array are not written. 
Do the above modifications to the input array in place and do not return anything.


## Объяснение
Задача предлагает вам изменять массив целых чисел, выполняя следующие действия:
1. Пройдите по массиву слева направо.
2. Если вы находитесь на числе, равном 0, удалите его и вставьте два нуля рядом с ним.
   
Обратите внимание, что длина массива не должна измениться после этого. 
Некоторые элементы могут быть удалены или перемещены вправо, чтобы освободить место для новых нулей.

Например, если вы получаете массив [1,0,2,3,0,4,5,0], после преобразования он должен стать [1,0,0,2,3,0,0,4].

Обратите внимание, что вы не должны использовать дополнительный массив для этой задачи. 
Вам нужно изменить входной массив на месте с сохранением его длины.


## Example 1:
````
Input: arr = [1,0,2,3,0,4,5,0]
Output: [1,0,0,2,3,0,0,4]
Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]
````


## Example 2:
````
Input: arr = [1,2,3]
Output: [1,2,3]
Explanation: After calling your function, the input array is modified to: [1,2,3]
````


## Constraints:
- 1 <= arr.length <= 10^4
- 0 <= arr[i] <= 9