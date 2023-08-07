/*
182. Duplicate Emails
Level - easy

Write an SQL query to report all the duplicate emails.
Note that it's guaranteed that the email field is not NULL.

Return the result table in any order.

The query result format is in the following example.


Example 1:

Input:
Person table:
+----+---------+
| id | email   |
+----+---------+
| 1  | a@b.com |
| 2  | c@d.com |
| 3  | a@b.com |
+----+---------+
Output:
+---------+
| Email   |
+---------+
| a@b.com |
+---------+
Explanation: a@b.com is repeated two times.
 */
Create table If Not Exists Person
(
    id int,
    email varchar(255)
);

Truncate table Person;
insert into Person (id, email) values ('1', 'a@b.com');
insert into Person (id, email) values ('2', 'c@d.com');
insert into Person (id, email) values ('3', 'a@b.com');


-- METHOD-1
select distinct p1.email
from person p1, person p2
where p1.id <> p2.id
  and p1.email = p2.email;


-- METHOD-2
select email as Email from Person group by email having count(email)>1;


-- METHOD-3
SELECT email AS Email
FROM (SELECT COUNT(email) AS EmailCount, email
      FROM Person
      GROUP BY email) AS Temp
WHERE EmailCount > 1;