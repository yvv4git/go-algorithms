/*
# 181. Employees Earning More Than Their Managers

Table: Employee

+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| name        | varchar |
| salary      | int     |
| managerId   | int     |
+-------------+---------+
id is the primary key (column with unique values) for this table.
Each row of this table indicates the ID of an employee, their name, salary, and the ID of their manager.

Write a solution to find the employees who earn more than their managers.

Return the result table in any order.

The result format is in the following example.


Example 1:

Input: 
Employee table:
+----+-------+--------+-----------+
| id | name  | salary | managerId |
+----+-------+--------+-----------+
| 1  | Joe   | 70000  | 3         |
| 2  | Henry | 80000  | 4         |
| 3  | Sam   | 60000  | Null      |
| 4  | Max   | 90000  | Null      |
+----+-------+--------+-----------+
Output: 
+----------+
| Employee |
+----------+
| Joe      |
+----------+
Explanation: Joe is the only employee who earns more than his manager.

*/

/*

Задача предполагает, что у нас есть таблица сотрудников, 
где каждый сотрудник имеет уникальный идентификатор (Id), имя (Name) 
и идентификатор его менеджера (ManagerId). Менеджер — это тоже сотрудник, который может управлять другими сотрудниками.

Необходимо найти всех сотрудников, которые зарабатывают больше, чем их непосредственные менеджеры. 
Другими словами, нужно выбрать тех сотрудников, у которых зарплата (Salary) выше, чем у их менеджеров.

Пример:

Предположим, у нас есть следующая таблица Employee:

Id	Name	Salary	ManagerId
1	Joe	70000	3
2	Henry	80000	4
3	Sam	60000	NULL
4	Max	90000	NULL

В этом примере:

Joe зарабатывает 70000, а его менеджер (Sam) зарабатывает 60000. Joe зарабатывает больше, чем его менеджер.
Henry зарабатывает 80000, а его менеджер (Max) зарабатывает 90000. Henry зарабатывает меньше, чем его менеджер.
*/

-- Создаем таблицу
CREATE TABLE Employee (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    salary INT,
    managerId INT
);

-- Наполнение таблицы данными
INSERT INTO Employee (id, name, salary, managerId) VALUES
(1, 'Joe', 70000, 3),
(2, 'Henry', 80000, 4),
(3, 'Sam', 60000, NULL),
(4, 'Max', 90000, NULL);

-- Решение
SELECT e1.name AS Employee
FROM Employee e1
JOIN Employee e2 ON e1.managerId = e2.id
WHERE e1.salary > e2.salary;