package migrations
-- Create the employee table
CREATE TABLE IF NOT EXISTS Employee (
                                        empid SERIAL PRIMARY KEY,
                                        name VARCHAR(100) NOT NULL,
    dob DATE NOT NULL,
    department VARCHAR(50) NOT NULL,
    salary DECIMAL(10, 2) NOT NULL,
    bossId INT,
    FOREIGN KEY (bossId) REFERENCES Employee(empid)
    );