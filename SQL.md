# Problems
[182. Duplicate Emails](https://leetcode.com/problems/duplicate-emails/description/)

[196. Delete Duplicate Emails](https://leetcode.com/problems/delete-duplicate-emails/description/)
# Contents
## 182. Duplicate Emails
### Description
Write a SQL query to find all duplicate emails in a table named Person.
```
+----+---------+
| Id | Email   |
+----+---------+
| 1  | a@b.com |
| 2  | c@d.com |
| 3  | a@b.com |
+----+---------+
```
For example, your query should return the following for the above table:
```
+---------+
| Email   |
+---------+
| a@b.com |
+---------+
```
Note: All emails are in lowercase.
### Solution
```sql
SELECT
    Email
FROM
    Person
GROUP BY
    Email
HAVING
    COUNT( * ) >= 2;
```
## 196. Delete Duplicate Emails
### Description

Write a SQL query to delete all duplicate email entries in a table named Person, keeping only unique emails based on its smallest Id.
```
+----+------------------+
| Id | Email            |
+----+------------------+
| 1  | john@example.com |
| 2  | bob@example.com  |
| 3  | john@example.com |
+----+------------------+
```
Id is the primary key column for this table.
For example, after running your query, the above Person table should have the following rows:
```
+----+------------------+
| Id | Email            |
+----+------------------+
| 1  | john@example.com |
| 2  | bob@example.com  |
+----+------------------+
```
Note:

Your output is the whole Person table after executing your sql. Use delete statement.

### Solution

```sql
DELETE p1
FROM
    Person p1,
    Person p2
WHERE
    p1.Email = p2.Email
    AND p1.Id > p2.Id
```
