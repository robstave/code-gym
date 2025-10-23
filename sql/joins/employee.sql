-- Sample table
-- Employee
-- | id | name  | salary | managerId |
-- | -- | ----- | ------ | --------- |
-- | 1  | Joe   | 70000  | 3         |
-- | 2  | Henry | 80000  | 4         |
-- | 3  | Sam   | 60000  | NULL      |
-- | 4  | Max   | 90000  | NULL      |

-- Problem: find employees who earn more than their managers.
-- We need to compare rows within the same table. The standard approach is a self-join
-- (join the table to itself) so we can refer to the manager and the employee as
-- separate aliases.

-- Quick explanation of the approach:
-- 1) Use two aliases for the table: m = manager, e = employee.
-- 2) Join on m.id = e.managerId so each employee row (e) is paired with their manager (m).
-- 3) Filter to rows where the employee's salary is greater than the manager's salary.
-- 4) Select the employee name and alias it as "Employee" to match the requested output.

-- Final solution:
SELECT e.name AS Employee
FROM employee m
JOIN employee e ON m.id = e.managerId
WHERE e.salary > m.salary;

-- Notes:
-- - Employees without a manager (managerId IS NULL) are automatically excluded because
--   the JOIN requires a matching manager row.
-- - The result for the sample data will be:
--     | Employee |
--     | -------- |
--     | Joe      |
--   (Joe earns 70000 while his manager Sam earns 60000.)