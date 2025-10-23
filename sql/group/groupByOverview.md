# GROUP BY and HAVING — a practical guide

This note shows how to:
- Aggregate rows with GROUP BY
- Filter groups with HAVING (and how it differs from WHERE)
- Use CASE to create derived groupings or bands
- Common pitfalls and tips

---

## Quick refresher

- GROUP BY collapses rows that share the same grouping keys and lets you compute aggregates like SUM, COUNT, AVG, MIN, MAX.
- WHERE filters individual rows before grouping.
- HAVING filters groups after the aggregate is computed.

Order of operations (simplified):
1) FROM / JOIN → 2) WHERE → 3) GROUP BY → 4) HAVING → 5) SELECT → 6) ORDER BY

Rules to remember:
- Every non-aggregated column in SELECT must appear in GROUP BY.
- COUNT(*) counts rows; COUNT(col) skips NULLs in col.
- HAVING can be used without GROUP BY to filter on a whole-table aggregate.

---

## Example scenario

Imagine a simple `sales` table with salesperson, region, quarter, and amount. We’ll use it below.

### 1) Total sales by region, keep only big regions
```sql
SELECT region, SUM(amount) AS total_sales
FROM sales
GROUP BY region
HAVING SUM(amount) > 20000
ORDER BY total_sales DESC;
```
- WHERE would filter rows before the SUM; HAVING uses the aggregated SUM to keep or discard groups.

### 2) Average sale per salesperson, but only if they made at least 2 sales
```sql
SELECT salesperson, AVG(amount) AS avg_sale, COUNT(*) AS orders
FROM sales
GROUP BY salesperson
HAVING COUNT(*) >= 2
ORDER BY avg_sale DESC;
```

### 3) Using CASE to create bands, then grouping and filtering on them
```sql
SELECT
  region,
  CASE
    WHEN amount >= 10000 THEN 'Large'
    WHEN amount >=  5000 THEN 'Medium'
    ELSE 'Small'
  END AS band,
  COUNT(*) AS orders,
  SUM(amount) AS total
FROM sales
GROUP BY region, band
HAVING band <> 'Small' -- keep only Medium and Large groups
ORDER BY region, band;
```
- CASE runs per-row; the result can be used as a grouping key.
- The alias `band` is valid in SELECT, ORDER BY, and HAVING in many dialects; if your dialect complains, repeat the CASE in HAVING.

### 4) HAVING without GROUP BY (whole table aggregate check)
```sql
SELECT SUM(amount) AS total_sales
FROM sales
HAVING SUM(amount) > 0; -- returns a row only if the total is > 0
```

---

## Other relevant tips
- Nulls form their own group. Use COALESCE to rename: `COALESCE(region, 'Unknown')`.
- Prefer WHERE for row-level filters (faster), HAVING for aggregate-level filters.
- You can combine multiple aggregates in HAVING, e.g., `HAVING SUM(amount) > 20000 AND COUNT(*) >= 3`.
- If you need distinct counts per group: `COUNT(DISTINCT customer_id)` (supported by most dialects).

---

## dbFiddle-ready script

Paste this into dbfiddle (PostgreSQL/MySQL/SQLite friendly). It creates the table, inserts sample data, and runs the queries above plus a couple of labeled headers to make outputs easier to scan.

```sql
-- Clean start (safe in most dialects; ignore errors if DROP fails)
DROP TABLE IF EXISTS sales;

CREATE TABLE sales (
  id        INTEGER,
  salesperson VARCHAR(50),
  region      VARCHAR(20),
  quarter     INTEGER,
  amount      INTEGER
);

INSERT INTO sales (id, salesperson, region, quarter, amount) VALUES
(1,  'Alice', 'North', 1,  7000),
(2,  'Alice', 'North', 2, 12000),
(3,  'Bob',   'North', 1,  4000),
(4,  'Bob',   'East',  2,  6000),
(5,  'Cara',  'East',  1, 15000),
(6,  'Cara',  'East',  3,  3000),
(7,  'Derek', 'West',  1,  8000),
(8,  'Derek', 'West',  2,  9000),
(9,  'Eva',   'South', 2,  2000),
(10, 'Eva',   'South', 3, 11000),
(11, 'Finn',  'South', 3,  5000),
(12, 'Gabe',   NULL,   1,  7000);

-- Show the raw data first (nice for dbfiddle)
SELECT 'sales (raw data)' AS section;
SELECT * FROM sales ORDER BY region NULLS LAST, salesperson, quarter;

-- 1) Total sales by region, keep only big regions
SELECT 'Total sales by region, HAVING SUM(amount) > 20000' AS section;
SELECT COALESCE(region, 'Unknown') AS region, SUM(amount) AS total_sales
FROM sales
GROUP BY region
HAVING SUM(amount) > 20000
ORDER BY total_sales DESC;

-- 2) Average sale per salesperson, only if they made at least 2 sales
SELECT 'Average sale per salesperson, HAVING COUNT(*) >= 2' AS section;
SELECT salesperson, AVG(amount) AS avg_sale, COUNT(*) AS orders
FROM sales
GROUP BY salesperson
HAVING COUNT(*) >= 2
ORDER BY avg_sale DESC;

-- 3) CASE bands per region (keeping Medium and Large only)
SELECT 'CASE banding by region (exclude Small via HAVING)' AS section;
SELECT
  COALESCE(region, 'Unknown') AS region,
  CASE
    WHEN amount >= 10000 THEN 'Large'
    WHEN amount >=  5000 THEN 'Medium'
    ELSE 'Small'
  END AS band,
  COUNT(*) AS orders,
  SUM(amount) AS total
FROM sales
GROUP BY region, band
HAVING band <> 'Small'
ORDER BY region, band;

-- 4) HAVING without GROUP BY: overall total check
SELECT 'Overall total via HAVING (no GROUP BY)' AS section;
SELECT SUM(amount) AS total_sales
FROM sales
HAVING SUM(amount) > 0;
```
