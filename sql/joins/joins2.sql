-- Clean start
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS teams;

-- Two simple tables
CREATE TABLE teams (
  team_id INT PRIMARY KEY,
  team_name VARCHAR(80) NOT NULL
);

CREATE TABLE users (
  user_id INT PRIMARY KEY,
  user_name VARCHAR(80) NOT NULL,
  team_id INT,        -- may be NULL or unmatched
  salary INT NOT NULL
);

-- Seed data
INSERT INTO teams (team_id, team_name) VALUES
(1, 'Red'),
(2, 'Blue'),
(3, 'Green'),
(4, 'Ghost');   -- no users on this team (to demo outer joins)

INSERT INTO users (user_id, user_name, team_id, salary) VALUES
(10, 'Alice',  1, 120000),
(11, 'Bob',    1,  90000),
(12, 'Cara',   2, 105000),
(13, 'Derek',  2,  70000),
(14, 'Eva',    3, 130000),
(15, 'Finn',   NULL, 65000),  -- no team
(16, 'Gabe',   99,  95000);   -- unmatched team_id

-- Common Join example + pretty-print table contents for dbfiddle
-- (Insert this near the top so dbfiddle shows table data first)

-- Show teams table with a label column to make output self-describing
SELECT 'teams' AS source, team_id, team_name FROM teams ORDER BY team_id;

-- Show users table with a label column
SELECT 'users' AS source, user_id, user_name, team_id, salary FROM users ORDER BY user_id;

-- Simple, common join: list users with their team name (if any)
SELECT 'Common join: users with team name (LEFT JOIN)' AS section;


SELECT u.user_id, u.user_name, COALESCE(t.team_name, '—') AS team_name
FROM users u
LEFT JOIN teams t ON u.team_id = t.team_id
ORDER BY u.user_id;

--  ================================================
-- 1) INNER JOIN: users that have a matching team
SELECT '1) INNER JOIN: users that have a matching team' AS section;
-- =========================================================



SELECT u.user_id, u.user_name, t.team_name
FROM users u
INNER JOIN teams t ON u.team_id = t.team_id
ORDER BY t.team_name, u.user_name;

-- =========================================================
-- 2) LEFT JOIN: keep all users, show team when present
--    (highlights NULL and unmatched team_id)

-- =========================================================



SELECT
  u.user_id, u.user_name, u.team_id,
  t.team_name,
  CASE
    WHEN t.team_id IS NULL THEN 'No matching team'
    ELSE 'Matched'
  END AS match_status
FROM users u
LEFT JOIN teams t ON u.team_id = t.team_id
ORDER BY match_status DESC, u.user_name;

-- =========================================================
-- 3) RIGHT JOIN: keep all teams, even if nobody is on them
-- =========================================================
SELECT '3) RIGHT JOIN: keep all teams, even if nobody is on them' AS section;


SELECT
  t.team_id, t.team_name,
  u.user_id, u.user_name
FROM users u
RIGHT JOIN teams t ON u.team_id = t.team_id
ORDER BY t.team_id, u.user_name;

-- =========================================================
-- 4) FULL OUTER JOIN (emulated in MySQL with UNION)
--    Show everything from both sides
SELECT '4) FULL OUTER JOIN (emulated with UNION): show everything from both sides' AS section;

-- =========================================================
-- Left side (all users)
SELECT
  COALESCE(u.user_id, -1) AS user_id,
  u.user_name,
  u.team_id AS user_team_id,
  t.team_id AS team_team_id,
  t.team_name
FROM users u
LEFT JOIN teams t ON u.team_id = t.team_id

UNION

-- Right side (all teams)
SELECT
  COALESCE(u.user_id, -1) AS user_id,
  u.user_name,
  u.team_id AS user_team_id,
  t.team_id AS team_team_id,
  t.team_name
FROM users u
RIGHT JOIN teams t ON u.team_id = t.team_id

ORDER BY team_team_id, user_name;

-- =========================================================
-- 5) CROSS JOIN: Cartesian product (use sparingly!)
--    (example: show every user paired with every team)
SELECT '5) CROSS JOIN: Cartesian product (limited output)' AS section;

-- =========================================================
SELECT u.user_name, t.team_name
FROM users u
CROSS JOIN teams t
ORDER BY u.user_name, t.team_name
LIMIT 12;

-- =========================================================
-- 6) IMPLICIT (old-style) JOIN: same result as inner join
SELECT '6) IMPLICIT JOIN (old-style): same as INNER JOIN' AS section;


-- =========================================================
SELECT u.user_id, u.user_name, t.team_name
FROM users u, teams t
WHERE u.team_id = t.team_id
ORDER BY t.team_name, u.user_name;

-- =========================================================
-- 7) USING(...) variant (safer than NATURAL JOIN)
SELECT '7) USING(...) variant (safer than NATURAL JOIN)' AS section;

-- =========================================================
SELECT u.user_id, u.user_name, t.team_name
FROM users u
JOIN teams t USING (team_id)
ORDER BY t.team_name, u.user_name;

-- =========================================================
-- 8) ANTI-JOIN (users with NO matching team)
--    (LEFT JOIN + WHERE IS NULL)
SELECT '8) ANTI-JOIN: users with NO matching team' AS section;

-- =========================================================
SELECT u.user_id, u.user_name, u.team_id
FROM users u
LEFT JOIN teams t ON u.team_id = t.team_id
WHERE t.team_id IS NULL
ORDER BY u.user_name;

-- =========================================================
-- 9) SEMI-JOIN (users that DO have a matching team)
--    (EXISTS)
SELECT '9) SEMI-JOIN: users that DO have a matching team (EXISTS)' AS section;

-- =========================================================
SELECT u.user_id, u.user_name, u.team_id
FROM users u
WHERE EXISTS (
  SELECT 1 FROM teams t WHERE t.team_id = u.team_id
)
ORDER BY u.user_name;

-- =========================================================
-- 10) Aggregation with JOIN: avg salary per team
SELECT '10) Aggregation with JOIN: avg salary per team' AS section;

-- =========================================================
SELECT t.team_name, ROUND(AVG(u.salary)) AS avg_salary
FROM users u
JOIN teams t ON u.team_id = t.team_id
GROUP BY t.team_name
ORDER BY avg_salary DESC;

-- =========================================================
-- 11) CASE example: salary band per user (with team)
SELECT '11) CASE example: salary band per user (with team)' AS section;

-- =========================================================
SELECT
  u.user_name,
  COALESCE(t.team_name, '—') AS team_name,
  u.salary,
  CASE
    WHEN u.salary >= 120000 THEN 'Platinum'
    WHEN u.salary >= 100000 THEN 'Gold'
    WHEN u.salary >=  80000 THEN 'Silver'
    ELSE 'Bronze'
  END AS salary_band
FROM users u
LEFT JOIN teams t ON u.team_id = t.team_id
ORDER BY salary_band DESC, u.salary DESC, u.user_name;

-- =========================================================
-- 12) CASE + GROUP: count users by band for each team
SELECT '12) CASE + GROUP: count users by band for each team' AS section;

-- =========================================================
WITH bands AS (
  SELECT
    COALESCE(t.team_name, 'No Team') AS team_name,
    CASE
      WHEN u.salary >= 120000 THEN 'Platinum'
      WHEN u.salary >= 100000 THEN 'Gold'
      WHEN u.salary >=  80000 THEN 'Silver'
      ELSE 'Bronze'
    END AS band
  FROM users u
  LEFT JOIN teams t ON u.team_id = t.team_id
)
SELECT team_name, band, COUNT(*) AS people
FROM bands
GROUP BY team_name, band
ORDER BY team_name, band;

-- =========================================================
-- 13) DISTINCT vs. COUNT(*) with JOIN
--     How many unique teams have at least one user?
SELECT '13) DISTINCT vs. COUNT(*) with JOIN: how many unique teams have at least one user?' AS section;

-- =========================================================
SELECT COUNT(DISTINCT t.team_id) AS teams_with_users
FROM users u
JOIN teams t ON u.team_id = t.team_id;

-- =========================================================
-- 14) Filtering before vs. after join (illustration)
--     Only join users earning >= 100k
SELECT '14) Filtering before vs. after join (only salary >= 100k)' AS section;

-- =========================================================
SELECT u.user_name, u.salary, t.team_name
FROM (SELECT * FROM users WHERE salary >= 100000) AS u
LEFT JOIN teams t ON u.team_id = t.team_id
ORDER BY u.salary DESC, u.user_name;


