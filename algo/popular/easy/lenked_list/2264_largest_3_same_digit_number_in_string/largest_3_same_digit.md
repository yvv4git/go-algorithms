# 2264. Largest 3-Same-Digit Number in String


## Level - easy

## Task 
You are given a string num representing a large integer.   
An integer is good if it meets the following conditions:  
- It is a substring of num with length 3.
- It consists of only one unique digit.

Return the maximum good integer as a string or an empty string "" if no such integer exists.  
Note:
- A substring is a contiguous sequence of characters within a string.
- There may be leading zeroes in num or a good integer.


## Объяснение
В этой задаче необходимо реализовать функцию largestGoodInteger, которая должна находить наибольшее число, 
состоящее из одной и той же цифры, длиной 3 символа.

## Example 1:
````
Input: num = "6777133339"
Output: "777"
Explanation: There are two distinct good integers: "777" and "333".
"777" is the largest, so we return "777".
````

## Example 2:
````
Input: num = "2300019"
Output: "000"
Explanation: "000" is the only good integer.
````

## Constraints:
- 3 <= num.length <= 1000
- num only consists of digits.