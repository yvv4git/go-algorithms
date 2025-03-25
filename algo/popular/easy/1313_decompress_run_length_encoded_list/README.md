# 1313. Decompress Run-Length Encoded List


## Level - easy


## Task
We are given a list nums of integers representing a list compressed with run-length encoding.

Consider each adjacent pair of elements [freq, val] = [nums[2*i], nums[2*i+1]] (with i >= 0).  
For each such pair, there are freq elements with value val concatenated in a sublist. 
Concatenate all the sublists from left to right to generate the decompressed list.

Return the decompressed list.


## Объяснение
На вход подаётся список (массив) целых чисел nums. 
Этот список представляет собой данные, закодированные по принципу RLE (Run-Length Encoding). 
Нужно декодировать этот список и вернуть разжатый вариант.

В RLE последовательность одинаковых элементов заменяется парой [freq, val], где freq — количество повторений, а val — значение. 
В данной задаче список nums состоит из нескольких таких пар, идущих подряд.

Пример:

Если nums = [1, 2, 3, 4], то это означает две пары:
- Первая пара: [1, 2] → число 2 повторяется 1 раз → [2].
- Вторая пара: [3, 4] → число 4 повторяется 3 раза → [4, 4, 4].

Объединив эти две части, получаем разжатый список [2, 4, 4, 4].


## Example 1:
```
Input: nums = [1,2,3,4]
Output: [2,4,4,4]
Explanation: The first pair [1,2] means we have freq = 1 and val = 2 so we generate the array [2].
The second pair [3,4] means we have freq = 3 and val = 4 so we generate [4,4,4].
At the end the concatenation [2] + [4,4,4] is [2,4,4,4].
```

## Example 2:
```
Input: nums = [1,1,2,3]
Output: [1,3,3]
```

## Constraints:
- 2 <= nums.length <= 100
- nums.length % 2 == 0
- 1 <= nums[i] <= 100