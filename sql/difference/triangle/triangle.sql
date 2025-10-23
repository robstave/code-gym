

-- https://www.db-fiddle.com/f/smGaSnERKgdLuM2M7LaYeJ/1


-- Create the TRIANGLES table
CREATE TABLE TRIANGLES (
    A INT,
    B INT,
    C INT
);

-- Insert sample data
INSERT INTO TRIANGLES (A, B, C) VALUES
(2, 2, 2),   -- Equilateral
(2, 3, 2),   -- Isosceles
(3, 4, 5),   -- Scalene
(1, 2, 3),   -- Not A Triangle
(10, 10, 15),-- Isosceles
(5, 7, 10);  -- Scalene

-- Query to classify the type of triangle
SELECT
  A, B, C,
  CASE
    WHEN A + B <= C OR A + C <= B OR B + C <= A THEN 'Not A Triangle'
    WHEN A = B AND B = C THEN 'Equilateral'
    WHEN A = B OR B = C OR A = C THEN 'Isosceles'
    ELSE 'Scalene'
  END AS TriangleType
FROM TRIANGLES;