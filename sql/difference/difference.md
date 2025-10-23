

// generate the sql for the difference in counts between two queries

```sql
select   COUNT(*) - COUNT(DISTINCT city) from station  
where country = 'USA';
```

