# 860. Lemonade Change


## Level - easy


## Task 
At a lemonade stand, each lemonade costs $5. 
Customers are standing in a queue to buy from you and order one at a time (in the order specified by bills). 
Each customer will only buy one lemonade and pay with either a $5, $10, or $20 bill. 
You must provide the correct change to each customer so that the net transaction is that the customer pays $5.

Note that you do not have any change in hand at first.

Given an integer array bills where bills[i] is the bill the ith customer pays, 
return true if you can provide every customer with the correct change, or false otherwise.


## Объяснение
В ларьке с лимонадом продается лимонад по 5 долларов за стакан. Каждый клиент платит либо 5 долларов, либо 10 долларов. 
Если клиент платит 10 долларов, то продавец должен вернуть сдачу в размере 5 долларов. 
Если у продавца нет 5-долларовых купюр, то он не может вернуть сдачу и клиент уходит без лимонада.

Дан массив bills, где bills[i] - это сумма, которую заплатил i-й клиент. 
Нужно определить, может ли продавец вернуть сдачу всем клиентам.

Примеры
- bills = [5, 5, 5, 10, 20]: Продавец может вернуть сдачу всем клиентам.
- bills = [5, 5, 10, 10]: Продавец не может вернуть сдачу последнему клиенту.


## Example 1:
````
Input: bills = [5,5,5,10,20]
Output: true
Explanation:
From the first 3 customers, we collect three $5 bills in order.
From the fourth customer, we collect a $10 bill and give back a $5.
From the fifth customer, we give a $10 bill and a $5 bill.
Since all customers got correct change, we output true.
````


## Example 2:
````
Input: bills = [5,5,10,10,20]
Output: false
Explanation:
From the first two customers in order, we collect two $5 bills.
For the next two customers in order, we collect a $10 bill and give back a $5 bill.
For the last customer, we can not give the change of $15 back because we only have two $10 bills.
Since not every customer received the correct change, the answer is false.
````


## Constraints:
- 1 <= bills.length <= 10^5
- bills[i] is either 5, 10, or 20.