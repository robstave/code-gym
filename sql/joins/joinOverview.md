# SQL Joins — a practical guide

This note shows how to:
- Pick the right join type (INNER, LEFT, RIGHT, FULL, CROSS)
- Write common patterns (SELF, SEMI, ANTI)
- Avoid pitfalls (duplication, NULLs, filters in the right place)
- Read join results quickly with small, labeled outputs

---

## Quick refresher

- INNER JOIN: keep rows that match on both sides.
- LEFT JOIN: keep all left rows; fill right side with NULLs when there’s no match.
- RIGHT JOIN: symmetric to LEFT JOIN (not in SQLite; flip tables or use LEFT).
- FULL OUTER JOIN: keep everything from both sides (not in MySQL/SQLite; use UNION of left/right).
- CROSS JOIN: cartesian product — every left row paired with every right row.

Order of operations (simplified):
1) FROM / JOIN → 2) WHERE → 3) GROUP BY → 4) HAVING → 5) SELECT → 6) ORDER BY

Tips:
- Always qualify ambiguous columns (u.team_id vs t.team_id).
- Filters in ON vs WHERE change results (especially with outer joins):
  - Put match logic in ON.
  - Put row-level filters in WHERE; but be careful with LEFT JOIN, a WHERE filter on right table can turn it into an INNER JOIN.
- Many-to-many can multiply rows — aggregate or DISTINCT thoughtfully.
- NULLs never match = two NULLs don’t join equal; use COALESCE if you want explicit groupings/labels.

---

## Example schema

We’ll use two small tables, teams and users. A user may have a NULL or unmatched team_id.

```
teams(team_id, team_name)
users(user_id, user_name, team_id, salary)
```

---

## Common joins by example

### 1) INNER JOIN — users that have a matching team
```sql
SELECT u.user_id, u.user_name, t.team_name
FROM users u
INNER JOIN teams t ON u.team_id = t.team_id
ORDER BY t.team_name, u.user_name;
```
- Drops users with NULL/unmatched team_id, and teams with no users.

### 2) LEFT JOIN — keep all users, annotate match status
```sql
SELECT
  u.user_id, u.user_name, u.team_id,
  t.team_name,
  CASE WHEN t.team_id IS NULL THEN 'No matching team' ELSE 'Matched' END AS match_status
FROM users u
LEFT JOIN teams t ON u.team_id = t.team_id
ORDER BY match_status DESC, u.user_name;
```
- RIGHT JOIN alternative: swap tables, or use RIGHT JOIN where supported.

### 3) FULL OUTER JOIN — everything from both sides (portable version)
```sql
-- MySQL/SQLite don’t support FULL OUTER JOIN; emulate with UNION
SELECT u.user_id, u.user_name, u.team_id AS user_team_id, t.team_id AS team_team_id, t.team_name
FROM users u LEFT JOIN teams t ON u.team_id = t.team_id
UNION
SELECT u.user_id, u.user_name, u.team_id AS user_team_id, t.team_id AS team_team_id, t.team_name
FROM users u RIGHT JOIN teams t ON u.team_id = t.team_id -- for SQLite flip to LEFT and swap sides
ORDER BY team_team_id, user_name;
```

### 4) CROSS JOIN — every user paired with every team
```sql
SELECT u.user_name, t.team_name
FROM users u
CROSS JOIN teams t
ORDER BY u.user_name, t.team_name
LIMIT 12; -- limit for readability
```

### 5) SELF JOIN — compare rows within the same table
Find employees who earn more than their managers (manager in same table):
```sql
SELECT e.name AS employee
FROM employee m
JOIN employee e ON m.id = e.managerId
WHERE e.salary > m.salary;
```

### 6) SEMI-JOIN (EXISTS) — keep left rows that have a match
```sql
SELECT u.user_id, u.user_name, u.team_id
FROM users u
WHERE EXISTS (SELECT 1 FROM teams t WHERE t.team_id = u.team_id)
ORDER BY u.user_name;
```

### 7) ANTI-JOIN — keep left rows with NO match
```sql
SELECT u.user_id, u.user_name, u.team_id
FROM users u
LEFT JOIN teams t ON u.team_id = t.team_id
WHERE t.team_id IS NULL
ORDER BY u.user_name;
```

### 8) USING(...) — shorthand when column names match
```sql
SELECT u.user_id, u.user_name, t.team_name
FROM users u
JOIN teams t USING (team_id)
ORDER BY t.team_name, u.user_name;
```

### 9) Aggregation with JOIN — average salary per team
```sql
SELECT t.team_name, ROUND(AVG(u.salary)) AS avg_salary
FROM users u
JOIN teams t ON u.team_id = t.team_id
GROUP BY t.team_name
ORDER BY avg_salary DESC;
```

---

## Other relevant notes
- Filtering before vs after join: push filters to WHERE or pre-filtered subqueries to reduce join size.
- DISTINCT vs GROUP BY: DISTINCT removes duplicates after projection; GROUP BY aggregates rows — they answer different questions.
- Beware of accidental INNER join behavior: applying a WHERE filter on the right-side columns after a LEFT JOIN can drop NULL-extended rows.
- For large joins, verify indexes on join keys (e.g., users.team_id, teams.team_id).

---

## dbFiddle-ready script

Paste this into dbfiddle (PostgreSQL/MySQL/SQLite). It creates the tables, inserts data, and runs labeled queries so outputs are easy to scan.

```sql
-- Clean start
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS teams;

CREATE TABLE teams (
  team_id INTEGER PRIMARY KEY,
  team_name VARCHAR(80) NOT NULL
);

CREATE TABLE users (
  user_id INTEGER PRIMARY KEY,
  user_name VARCHAR(80) NOT NULL,
  team_id INTEGER,        -- may be NULL or unmatched
  salary INTEGER NOT NULL
);

INSERT INTO teams (team_id, team_name) VALUES
(1, 'Red'),
(2, 'Blue'),
(3, 'Green'),
(4, 'Ghost');   -- no users on this team

INSERT INTO users (user_id, user_name, team_id, salary) VALUES
(10, 'Alice',  1, 120000),
(11, 'Bob',    1,  90000),
(12, 'Cara',   2, 105000),
(13, 'Derek',  2,  70000),
(14, 'Eva',    3, 130000),
(15, 'Finn',   NULL, 65000),  -- no team
(16, 'Gabe',   99,  95000);   -- unmatched team_id

-- Show raw data first
SELECT 'teams (raw data)' AS section;
SELECT * FROM teams ORDER BY team_id;
SELECT 'users (raw data)' AS section;
SELECT * FROM users ORDER BY user_id;

-- 1) INNER JOIN
SELECT '1) INNER JOIN: users with matching team' AS section;
SELECT u.user_id, u.user_name, t.team_name
FROM users u
INNER JOIN teams t ON u.team_id = t.team_id
ORDER BY t.team_name, u.user_name;

-- 2) LEFT JOIN
SELECT '2) LEFT JOIN: keep all users, label match' AS section;
SELECT u.user_id, u.user_name, u.team_id, COALESCE(t.team_name, '—') AS team_name,
       CASE WHEN t.team_id IS NULL THEN 'No match' ELSE 'Matched' END AS match_status
FROM users u
LEFT JOIN teams t ON u.team_id = t.team_id
ORDER BY match_status DESC, u.user_name;

-- 3) FULL OUTER JOIN (portable)
SELECT '3) FULL OUTER JOIN (via UNION)' AS section;
SELECT u.user_id, u.user_name, u.team_id AS user_team_id, t.team_id AS team_team_id, t.team_name
FROM users u LEFT JOIN teams t ON u.team_id = t.team_id
UNION
SELECT u.user_id, u.user_name, u.team_id AS user_team_id, t.team_id AS team_team_id, t.team_name
FROM users u RIGHT JOIN teams t ON u.team_id = t.team_id
ORDER BY team_team_id, user_name;

-- 4) CROSS JOIN (limited)
SELECT '4) CROSS JOIN (limited output)' AS section;
SELECT u.user_name, t.team_name
FROM users u
CROSS JOIN teams t
ORDER BY u.user_name, t.team_name
LIMIT 12;

-- 5) SEMI-JOIN (EXISTS)
SELECT '5) SEMI-JOIN: users that have a team' AS section;
SELECT u.user_id, u.user_name, u.team_id
FROM users u
WHERE EXISTS (SELECT 1 FROM teams t WHERE t.team_id = u.team_id)
ORDER BY u.user_name;

-- 6) ANTI-JOIN (no matching team)
SELECT '6) ANTI-JOIN: users with NO matching team' AS section;
SELECT u.user_id, u.user_name, u.team_id
FROM users u
LEFT JOIN teams t ON u.team_id = t.team_id
WHERE t.team_id IS NULL
ORDER BY u.user_name;

-- 7) USING(...)
SELECT '7) USING(team_id) variant' AS section;
SELECT u.user_id, u.user_name, t.team_name
FROM users u
JOIN teams t USING (team_id)
ORDER BY t.team_name, u.user_name;

-- 8) Aggregation with JOIN
SELECT '8) Aggregation: average salary per team' AS section;
SELECT t.team_name, ROUND(AVG(u.salary)) AS avg_salary
FROM users u
JOIN teams t ON u.team_id = t.team_id
GROUP BY t.team_name
ORDER BY avg_salary DESC;
```

Dialect notes:
- RIGHT JOIN isn’t available in SQLite. If your dbfiddle uses SQLite, replace the RIGHT JOIN half of the FULL OUTER JOIN emulation with another LEFT JOIN by swapping table order, e.g.:
  - `SELECT ... FROM teams t LEFT JOIN users u ON u.team_id = t.team_id`.
- FULL OUTER JOIN isn’t available in MySQL/SQLite; the UNION pattern above is a common workaround.
