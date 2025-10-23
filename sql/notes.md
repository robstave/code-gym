

# SQL Joins Examples

## Left Join
A LEFT JOIN returns all records from the left table (table1), and the matched records from the right table (table2). The result is NULL from the right side, if there is no match.

```sql
 

SELECT p.firstName, p.lastname,   a.city, a.state 
FROM Person p 
LEFT JOIN address a ON  
  a.personId = p.personId;
```




----
 where x IS null


 SELECT Name FROM Products
WHERE ProductId IN
( SELECT ProductId 
  FROM ProductSales
  WHERE ProductSales.Date BETWEEN  Y AND Z 
  GROUP BY ProductId
  HAVING SUM(ProductSales.Qty) > x
)

