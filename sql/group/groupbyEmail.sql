

https://www.db-fiddle.com/f/ggKTQQw162mQDN9KuTPUUg/0
https://www.db-fiddle.com/f/ggKTQQw162mQDN9KuTPUUg/1


CREATE TABLE Person (
  id INT,
  email VARCHAR(50)
);

INSERT INTO Person (id, email) VALUES
(1, 'a@b.com'),
(2, 'c@d.com'),
(3, 'a@b.com');


SELECT email, COUNT(*) AS count_per_email
FROM Person         
GROUP BY email;


-- Expected output:-- email    | count_per_email
-- ---------+-----------------
-- a@b.com | 2
-- c@d.com | 1

-- find the emails that occur more than once
SELECT email    
FROM Person
GROUP BY email
HAVING COUNT(*) > 1;
-- Expected output:-- email
-- ----------
-- a@b.com      


-- show counts
SELECT email, count(*) as counts
FROM Person
GROUP BY email ;


-- Find duplicate emails
select 'find dups' AS '';

-- show emails that have dups
SELECT email
FROM Person
GROUP BY email
HAVING COUNT(*) > 1;


-- Show dups with higher ids
-- note the cross join produces a lot of stuff

SELECT p1.id, p1.email FROM Person p1
JOIN Person p2 
ON p1.email = p2.email AND p1.id > p2.id ;


DELETE p1 FROM Person p1
JOIN Person p2 
ON p1.email = p2.email AND p1.id > p2.id;

Select * from Person;