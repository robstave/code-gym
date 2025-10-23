https://www.db-fiddle.com/f/rQY81RkXMi5WSJA6SAjFKR/0

-- Create the EMPLOYEE table
CREATE TABLE EMPLOYEE (
    ID INT PRIMARY KEY,
    NAME VARCHAR(50),
    MONTHLY_SALARY INT,
    MONTHS_WORKED INT
);

-- Insert sample data
INSERT INTO EMPLOYEE (ID, NAME, MONTHLY_SALARY, MONTHS_WORKED) VALUES
(1, 'Alice', 1000, 5),
(2, 'Bob', 1200, 4),
(3, 'Carol', 1000, 5),
(4, 'David', 900, 6),
(5, 'Eve', 800, 7);

-- Query: Find max total earnings and how many employees have it
SELECT
    (MONTHLY_SALARY * MONTHS_WORKED) AS total_earnings,
    COUNT(*) AS employee_count
FROM EMPLOYEE
WHERE (MONTHLY_SALARY * MONTHS_WORKED) = (
    SELECT MAX(MONTHLY_SALARY * MONTHS_WORKED)
    FROM EMPLOYEE
)
GROUP BY total_earnings;


note..you cant use total_earnings in the WHERE clause directly
 because SQL does not allow alias usage in that context.
