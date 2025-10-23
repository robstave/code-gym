

```sql
SELECT * FROM patients 
WHERE first_name in (SELECT first_name FROM doctors)
```



```sql
SELECT * FROM patients
  WHERE weight BETWEEN 100 AND 120
  AND province_id NOT IN ('ON', 'SK', 'AB');
```