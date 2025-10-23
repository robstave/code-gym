https://www.db-fiddle.com/f/VnMTKNLxS7Me5AvNkdsyr/0

-- dbfiddle.sql
-- Schema + sample data + queries to demonstrate selecting cities that start with vowels
-- Paste this into https://dbfiddle.uk (choose PostgreSQL, MySQL, or SQLite; queries use standard SQL compatible with all)

DROP TABLE IF EXISTS station;

CREATE TABLE station (
  id SERIAL PRIMARY KEY,
  city TEXT NOT NULL
);

INSERT INTO station (city) VALUES
('alphaville'),
('Boston'),
('orlando'),
('Eden'),
('Irvine'),
('umbra'),
('Zurich'),
('appleton'),
('Oakland'),
('edinburgh'),
('istanbul'),
('osaka'),
('Anchorage'),
('quebec');

-- 1) Case-sensitive: cities starting with lowercase vowels (a,e,i,o,u)
SELECT DISTINCT city
FROM station
WHERE city LIKE 'a%' OR city LIKE 'e%' OR city LIKE 'i%' OR city LIKE 'o%' OR city LIKE 'u%'
ORDER BY city;

-- 2) Case-insensitive (PostgreSQL ILIKE; MySQL: use LOWER(city) LIKE ... )
-- PostgreSQL
SELECT DISTINCT city
FROM station
WHERE city ILIKE 'a%' OR city ILIKE 'e%' OR city ILIKE 'i%' OR city ILIKE 'o%' OR city ILIKE 'u%'
ORDER BY city;

-- 3) Using LOWER() to normalize (works in most dialects)
SELECT DISTINCT city
FROM station
WHERE LOWER(city) LIKE 'a%' OR LOWER(city) LIKE 'e%' OR LOWER(city) LIKE 'i%' OR LOWER(city) LIKE 'o%' OR LOWER(city) LIKE 'u%'
ORDER BY city;

-- 4) Using regex (Postgres) to match start with vowel (case-insensitive)
-- SELECT DISTINCT city FROM station WHERE city ~* '^[aeiou]' ORDER BY city;

-- 5) A compact IN-style variant using substring (works in many dialects)
SELECT DISTINCT city
FROM station
WHERE substring(city from 1 for 1) IN ('a','A','e','E','i','I','o','O','u','U')
ORDER BY city;

-- === Cities that END with a vowel ===
-- User-provided snippet (case-sensitive; matches lowercase endings):
SELECT DISTINCT city FROM station
WHERE city LIKE '%a' OR city LIKE '%e' OR city LIKE '%i' OR city LIKE '%o' OR city LIKE '%u';

-- Case-insensitive using LOWER() (portable across dialects):
SELECT DISTINCT city FROM station
WHERE LOWER(city) LIKE '%a' OR LOWER(city) LIKE '%e' OR LOWER(city) LIKE '%i' OR LOWER(city) LIKE '%o' OR LOWER(city) LIKE '%u'
ORDER BY city;

-- PostgreSQL regex, case-insensitive (compact):
-- SELECT DISTINCT city FROM station WHERE city ~* '[aeiou]$' ORDER BY city;

-- Using substring to check last character (portable):
SELECT DISTINCT city FROM station
WHERE substring(city from char_length(city) for 1) IN ('a','A','e','E','i','I','o','O','u','U')
ORDER BY city;

-- === Cities that DO NOT start with a vowel ===
-- User-provided snippet (case-sensitive; excludes lowercase vowel starts):
SELECT DISTINCT city FROM station
WHERE NOT (city LIKE 'a%' OR city LIKE 'e%' OR city LIKE 'i%' OR city LIKE 'o%' OR city LIKE 'u%');

-- Case-insensitive: use LOWER() to normalize (portable)
SELECT DISTINCT city FROM station
WHERE NOT (LOWER(city) LIKE 'a%' OR LOWER(city) LIKE 'e%' OR LOWER(city) LIKE 'i%' OR LOWER(city) LIKE 'o%' OR LOWER(city) LIKE 'u%')
ORDER BY city;

-- PostgreSQL regex alternative (case-insensitive):
-- SELECT DISTINCT city FROM station WHERE NOT (city ~* '^[aeiou]') ORDER BY city;








