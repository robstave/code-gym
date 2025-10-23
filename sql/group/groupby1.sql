DROP TABLE IF EXISTS sales;

CREATE TABLE sales (
  id INT PRIMARY KEY AUTO_INCREMENT,
  region VARCHAR(20),
  salesperson VARCHAR(30),
  product VARCHAR(20),
  qty INT,
  price DECIMAL(8,2)
);

INSERT INTO sales (region, salesperson, product, qty, price) VALUES
('East',  'Alice', 'Widget',  10, 20.00),
('East',  'Alice', 'Gizmo',    5, 35.00),
('East',  'Bob',   'Widget',  12, 20.00),
('West',  'Cara',  'Widget',  15, 25.00),
('West',  'Cara',  'Gizmo',    8, 40.00),
('West',  'Dan',   'Widget',   6, 25.00),
('North', 'Eli',   'Widget',   9, 22.00),
('South', 'Faye',  'Gizmo',   14, 38.00),
('South', 'Faye',  'Widget',   4, 22.00);

-- =========================================================
-- 1) Basic GROUP BY (one column)
-- =========================================================
SELECT region, SUM(qty) AS total_qty
FROM sales
GROUP BY region
ORDER BY total_qty DESC;

-- =========================================================
-- 2) GROUP BY multiple columns
-- =========================================================
SELECT region, salesperson, SUM(qty * price) AS total_sales
FROM sales
GROUP BY region, salesperson
ORDER BY region, total_sales DESC;

-- =========================================================
-- 3) GROUP BY expression (computed field)
-- =========================================================
SELECT
  region,
  CASE
    WHEN price < 25 THEN 'Low'
    WHEN price BETWEEN 25 AND 35 THEN 'Mid'
    ELSE 'High'
  END AS price_band,
  SUM(qty) AS items_sold
FROM sales
GROUP BY region, price_band
ORDER BY region, price_band;

-- =========================================================
-- 4) Using HAVING to filter aggregated results
-- =========================================================
SELECT salesperson, SUM(qty * price) AS revenue
FROM sales
GROUP BY salesperson
HAVING SUM(qty * price) > 400
ORDER BY revenue DESC;

-- =========================================================
-- 5) GROUP BY with COUNT(DISTINCT)
-- =========================================================
SELECT region, COUNT(DISTINCT product) AS distinct_products
FROM sales
GROUP BY region;

-- =========================================================
-- 6) GROUP BY with ROLLUP (MySQL extension)
-- =========================================================
SELECT region, SUM(qty * price) AS total_sales
FROM sales
GROUP BY region WITH ROLLUP;

-- =========================================================
-- 7) GROUP BY with aggregate in CASE (conditional totals)
-- =========================================================
SELECT
  region,
  SUM(CASE WHEN product = 'Widget' THEN qty ELSE 0 END) AS widget_qty,
  SUM(CASE WHEN product = 'Gizmo'  THEN qty ELSE 0 END) AS gizmo_qty
FROM sales
GROUP BY region;
