-- Create database
CREATE DATABASE IF NOT EXISTS banking;
USE banking;

-- Drop table if exists (optional, for re-run safety)
DROP TABLE IF EXISTS customers;

-- Create customers table
-- CREATE TABLE customers (
--     customer_id INT(11) NOT NULL,
--     name VARCHAR(100) NOT NULL,
--     date_of_birth DATE NOT NULL,
--     city VARCHAR(100) NOT NULL,
--     zipcode VARCHAR(10) NOT NULL,
--     status TINYINT(1) NOT NULL,
--     PRIMARY KEY (customer_id)
-- );

-- Insert sample data
-- INSERT INTO customers (customer_id, name, date_of_birth, city, zipcode, status) VALUES
-- (2000, 'Steve',  '1978-12-15', 'Delhi',          '110075', 1),
-- (2001, 'Arian',  '1988-05-21', 'Newburgh, NY',   '12550',  1),
-- (2002, 'Hadley', '1988-04-30', 'Englewood, NJ',  '07631',  1),
-- (2003, 'Ben',    '1988-01-04', 'Manchester, NH','03102',  0),
-- (2004, 'Nina',   '1988-05-14', 'Clarkston, MI',  '48348',  1),
-- (2005, 'Osman',  '1988-11-08', 'Hyattsville, MD','20782',  0);
