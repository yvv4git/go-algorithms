/*
# 183. Customers Who Never Order

Table: Customers

+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| name        | varchar |
+-------------+---------+
id is the primary key (column with unique values) for this table.
Each row of this table indicates the ID and name of a customer.


Table: Orders

+-------------+------+
| Column Name | Type |
+-------------+------+
| id          | int  |
| customerId  | int  |
+-------------+------+
id is the primary key (column with unique values) for this table.
customerId is a foreign key (reference columns) of the ID from the Customers table.
Each row of this table indicates the ID of an order and the ID of the customer who ordered it.


Write a solution to find all customers who never order anything.

Return the result table in any order.

The result format is in the following example.



## Example 1:

Input: 
Customers table:
+----+-------+
| id | name  |
+----+-------+
| 1  | Joe   |
| 2  | Henry |
| 3  | Sam   |
| 4  | Max   |
+----+-------+
Orders table:
+----+------------+
| id | customerId |
+----+------------+
| 1  | 3          |
| 2  | 1          |
+----+------------+
Output: 
+-----------+
| Customers |
+-----------+
| Henry     |
| Max       |
+-----------+

*/


/*
## ОБЪЯСНЕНИЕ
Нам нужно найти всех клиентов, которые никогда не делали заказов. 
Это можно сделать с помощью SQL-запроса, который будет использовать оператор LEFT JOIN для соединения таблиц Customers и Orders 
и затем отфильтрует клиентов, у которых нет соответствующих записей в таблице Orders.
*/

-- Prepare
-- Создание таблицы Customers
CREATE TABLE Customers (
    id INT PRIMARY KEY,
    name VARCHAR(255)
);

-- Создание таблицы Orders
CREATE TABLE Orders (
    id INT PRIMARY KEY,
    customerId INT,
    FOREIGN KEY (customerId) REFERENCES Customers(id)
);

-- Вставка данных в таблицу Customers
INSERT INTO Customers (id, name) VALUES
(1, 'Joe'),
(2, 'Henry'),
(3, 'Sam'),
(4, 'Max');

-- Вставка данных в таблицу Orders
INSERT INTO Orders (id, customerId) VALUES
(1, 3),
(2, 1);


-- Solution
SELECT c.name AS Customers
FROM Customers c
LEFT JOIN Orders o ON c.id = o.customerId
WHERE o.customerId IS NULL;