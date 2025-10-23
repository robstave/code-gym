https://www.db-fiddle.com/f/mzmBUKtFcb2d2W8xjeUtcT/0



CREATE TABLE Weather (
  id INT,
  recordDate DATE,
  temperature INT
);

INSERT INTO Weather (id, recordDate, temperature) VALUES
(1, '2015-01-01', 10),
(2, '2015-01-02', 25),
(3, '2015-01-03', 20),
(4, '2015-01-04', 30);

-- Find all dates with higher temperature than the previous date
SELECT w1.recordDate AS warmerDate
FROM Weather w1                 
JOIN Weather w2                 
ON DATEDIFF(w1.recordDate, w2.recordDate) = 1
WHERE w1.temperature > w2.temperature;

