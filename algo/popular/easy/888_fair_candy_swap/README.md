# 888. Fair Candy Swap


## Level - easy


## Task
Alice and Bob have a different total number of candies. 
You are given two integer arrays aliceSizes and bobSizes where aliceSizes[i] is the number of candies of the ith box of candy 
that Alice has and bobSizes[j] is the number of candies of the jth box of candy that Bob has.

Since they are friends, they would like to exchange one candy box each so that after the exchange, they both have the same total amount of candy. 
The total amount of candy a person has is the sum of the number of candies in each box they have.

Return an integer array answer where answer[0] is the number of candies in the box that Alice must exchange, 
and answer[1] is the number of candies in the box that Bob must exchange. If there are multiple answers, 
you may return any one of them. It is guaranteed that at least one answer exists.


## Объяснение
У вас есть два массива A и B, представляющие собой наборы конфет, которые имеют Алиса и Боб соответственно. К
аждый элемент массива представляет собой количество конфет определенного вида. 
Алиса и Боб хотят обменяться одной конфетой (один элемент из массива A на один элемент из массива B) так, 
чтобы после обмена сумма конфет у Алисы стала равна сумме конфет у Боба.

Вам нужно вернуть массив из двух элементов [x, y], где x — это конфета, 
которую Алиса отдаст, а y — конфета, которую Боб отдаст. 
Если возможных вариантов обмена несколько, можно вернуть любой из них.

Пример:
```
A = [1, 1]
B = [2, 2]
```
Выход:
```
[1, 2]
```

Объяснение:
- Сумма конфет у Алисы: 1 + 1 = 2
- Сумма конфет у Боба: 2 + 2 = 4
- Чтобы уравнять суммы, Алиса должна отдать конфету 1 и получить конфету 2 от Боба.


## Example 1:
```
Input: aliceSizes = [1,1], bobSizes = [2,2]
Output: [1,2]
```


## Example 2:
```
Input: aliceSizes = [1,2], bobSizes = [2,3]
Output: [1,2]
```


## Example 3:
```
Input: aliceSizes = [2], bobSizes = [1,3]
Output: [2,3]
```


## Constraints:
- 1 <= aliceSizes.length, bobSizes.length <= 10^4
- 1 <= aliceSizes[i], bobSizes[j] <= 10^5
- Alice and Bob have a different total number of candies.
- There will be at least one valid answer for the given input.