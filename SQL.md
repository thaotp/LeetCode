# Problems
[175. Combine Two Tables](175.-Combine-Two-Tables)

[182. Duplicate Emails](#182.-Duplicate-Emails)

[196. Delete Duplicate Emails](#196.-Delete-Duplicate-Emails)
# Contents
## 175. Combine Two Tables
### Description
https://leetcode.com/problems/combine-two-tables/description/

Table: Person
```
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| PersonId    | int     |
| FirstName   | varchar |
| LastName    | varchar |
+-------------+---------+
```
PersonId is the primary key column for this table.
Table: Address
```
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| AddressId   | int     |
| PersonId    | int     |
| City        | varchar |
| State       | varchar |
+-------------+---------+
```
AddressId is the primary key column for this table.
 

Write a SQL query for a report that provides the following information for each person in the Person table, regardless if there is an address for each of those people:
```
FirstName, LastName, City, State
```
### Solution
```SQL
# Write your MySQL query statement below
SELECT FirstName,
       LastName,
       City,
       State
FROM Person
LEFT JOIN Address ON Person.PersonId = Address.PersonId ;
```

## 182. Duplicate Emails
### Description
https://leetcode.com/problems/duplicate-emails/description/

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
https://leetcode.com/problems/delete-duplicate-emails/description/

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
