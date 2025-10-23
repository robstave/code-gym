

--https://www.db-fiddle.com/f/2FmY4V1o7ovBswjQa7f7EW/0
-- Create the STUDENTS table
CREATE TABLE STUDENTS (
    ID INT PRIMARY KEY,
    Name VARCHAR(50),
    Marks INT
);

-- Insert sample data
INSERT INTO STUDENTS (ID, Name, Marks) VALUES
(1, 'Bobby', 82),
(2, 'Robby', 91),
(3, 'Annie', 85),
(4, 'Danny', 70),
(5, 'Casey', 88),
(6, 'Jenny', 90),
(7, 'Mickey', 60);

-- Replace 75 with your desired minimum mark
SELECT Name
FROM STUDENTS
WHERE Marks > 75
ORDER BY RIGHT(Name, 3), ID ASC;