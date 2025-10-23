

-- https://www.db-fiddle.com/f/oEz6NbEgTvT1wenC8spNWz/0


-- === Clean start ===
DROP TABLE IF EXISTS PEOPLE;
DROP TABLE IF EXISTS AREAS;

CREATE TABLE AREAS (
  area_code INT PRIMARY KEY,
  city      VARCHAR(80) NOT NULL,
  state     CHAR(2) NOT NULL
);

CREATE TABLE PEOPLE (
  id         INT PRIMARY KEY,
  name       VARCHAR(80) NOT NULL,
  occupation VARCHAR(80) NOT NULL,
  area_code  INT,                -- allow NULLs / unmatched values for join demos
  income     INT NOT NULL
  -- (FK intentionally omitted so we can demo unmatched codes like 999)
);

-- === Seed data ===
INSERT INTO AREAS (area_code, city, state) VALUES
(212, 'New York', 'NY'),
(332, 'New York', 'NY'),
(646, 'New York', 'NY'),
(718, 'New York', 'NY'),
(917, 'New York', 'NY'),
(213, 'Los Angeles', 'CA'),
(310, 'Los Angeles', 'CA'),
(424, 'Los Angeles', 'CA'),
(562, 'Los Angeles', 'CA'),
(312, 'Chicago', 'IL'),
(773, 'Chicago', 'IL'),
(872, 'Chicago', 'IL'),
(305, 'Miami', 'FL'),
(786, 'Miami', 'FL'),
(415, 'San Francisco', 'CA'),
(628, 'San Francisco', 'CA'),
(619, 'San Diego', 'CA'),
(858, 'San Diego', 'CA'),
(714, 'Orange County', 'CA'),
(657, 'Orange County', 'CA'),
(214, 'Dallas', 'TX'),
(469, 'Dallas', 'TX'),
(972, 'Dallas', 'TX'),
(209, 'Stockton', 'CA'),
(702, 'Las Vegas', 'NV');

INSERT INTO PEOPLE (id, name, occupation, area_code, income) VALUES
( 1, 'Alice Johnson',      'Software Engineer', 415, 145000),
( 2, 'Bob Smith',          'Data Analyst',      628, 110000),
( 3, 'Carlos Reyes',       'Teacher',           212,  65000),
( 4, 'Diana Park',         'Nurse',             310,  98000),
( 5, 'Ethan Brown',        'Sales',             424,  87000),
( 6, 'Fiona Lee',          'Product Manager',   213, 155000),
( 7, 'George Miller',      'Lawyer',            212, 190000),
( 8, 'Hannah Kim',         'Graphic Designer',  718,  82000),
( 9, 'Ian Thompson',       'Electrician',       305,  72000),
(10, 'Jade Nguyen',        'Software Engineer', 619, 135000),
(11, 'Kevin Patel',        'Accountant',        773,  97000),
(12, 'Lily Chen',          'Nurse',             786, 102000),
(13, 'Marcus Davis',       'Sales',             312,  92000),
(14, 'Nora Garcia',        'Teacher',           714,  72000),
(15, 'Owen Scott',         'Software Engineer', 657, 128000),
(16, 'Priya Singh',        'Data Scientist',    858, 165000),
(17, 'Quinn Rogers',       'Barista',           917,  42000),
(18, 'Rosa Martinez',      'Chef',              415,  68000),
(19, 'Sam Wilson',         'Mechanic',          469,  75000),
(20, 'Tina Alvarez',       'HR Manager',        972, 105000),
(21, 'Uma Kapoor',         'Lawyer',            310, 200000),
(22, 'Victor Huang',       'Product Manager',   424, 158000),
(23, 'Wendy Brooks',       'Teacher',           NULL, 64000),  -- NULL
(24, 'Xavier Ortiz',       'Software Engineer', 999, 130000),  -- unmatched
(25, 'Yara Haddad',        'Nurse',             213,  99000),
(26, 'Zane Foster',        'Sales',             214,  88000);


-- === Basic CROSS JOIN (comma syntax) ===
SELECT
  p.name,
  a.city,
  a.state
FROM PEOPLE p, AREAS a
WHERE p.area_code = a.area_code;

-- same as above but using explicit J  OIN syntax
SELECT
  p.name,
  a.city,
  a.state
FROM PEOPLE p
JOIN AREAS a ON p.area_code = a.area_code;  





-- 1) INNER JOIN: only matched area codes
SELECT
  p.id, p.name, p.occupation, p.area_code, a.city, a.state, p.income
FROM PEOPLE p
JOIN AREAS a ON p.area_code = a.area_code
ORDER BY a.state, a.city, p.name;

-- 2) LEFT JOIN: keep everyone; flag unmatched
SELECT
  p.id, p.name, p.occupation, p.area_code, a.city, a.state, p.income,
  CASE WHEN a.area_code IS NULL THEN 'No matching area on file'
       ELSE 'Area matched' END AS area_match_status
FROM PEOPLE p
LEFT JOIN AREAS a ON p.area_code = a.area_code
ORDER BY area_match_status DESC, p.name;

-- 3) RIGHT JOIN: show all areas (even with no people)
SELECT
  a.area_code, a.city, a.state,
  p.id AS person_id, p.name AS person_name
FROM PEOPLE p
RIGHT JOIN AREAS a ON p.area_code = a.area_code
ORDER BY a.state, a.city, a.area_code, person_name;

-- 4) Aggregate with JOIN: avg income by city/state
SELECT
  a.city, a.state, AVG(p.income) AS avg_income
FROM PEOPLE p
JOIN AREAS a ON p.area_code = a.area_code
GROUP BY a.city, a.state
ORDER BY avg_income DESC;

-- 5) CASE: income bands per person (with location if available)
SELECT
  p.name,
  COALESCE(a.city, 'Unknown') AS city,
  COALESCE(a.state, '--') AS state,
  p.income,
  CASE
    WHEN p.income >= 170000 THEN 'Platinum'
    WHEN p.income >= 140000 THEN 'Gold'
    WHEN p.income >= 100000 THEN 'Silver'
    WHEN p.income >=  70000 THEN 'Bronze'
    ELSE 'Entry'
  END AS income_band
FROM PEOPLE p
LEFT JOIN AREAS a ON p.area_code = a.area_code
ORDER BY income_band DESC, p.income DESC, p.name;

-- 6) CASE + GROUP: occupation families per city
WITH people_labeled AS (
  SELECT
    p.*,
    CASE
      WHEN p.occupation IN ('Software Engineer','Data Scientist','Data Analyst','Product Manager') THEN 'Tech'
      WHEN p.occupation IN ('Teacher') THEN 'Education'
      WHEN p.occupation IN ('Nurse') THEN 'Healthcare'
      WHEN p.occupation IN ('Lawyer') THEN 'Legal'
      WHEN p.occupation IN ('Sales') THEN 'Sales'
      WHEN p.occupation IN ('Graphic Designer','Chef','Barista','Electrician','Mechanic','Accountant','HR Manager') THEN 'Other'
      ELSE 'Other'
    END AS occupation_family
  FROM PEOPLE p
)
SELECT
  a.city, a.state, pl.occupation_family,
  COUNT(*) AS people_count,
  ROUND(AVG(pl.income)) AS avg_income
FROM people_labeled pl
JOIN AREAS a ON pl.area_code = a.area_code
GROUP BY a.city, a.state, pl.occ
