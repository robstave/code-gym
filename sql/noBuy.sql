https://www.db-fiddle.com/f/jRjqfMmypanyFWXUyBKH22/0


given the following tables:

FIND THE CUSTOMERS WHO DID NOT MAKE ANY ORDERS.

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




--------
SELECT c.name AS Customers
FROM Customers c        
LEFT JOIN Orders o ON c.id = o.customerId
WHERE o.id IS NULL;

-- Alternative solution using NOT IN
-- SELECT name AS Customers
-- FROM Customers
-- WHERE id NOT IN (SELECT customerId FROM Orders);

