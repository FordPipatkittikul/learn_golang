-- Create database
CREATE DATABASE IF NOT EXISTS banking;
USE banking;

-- Drop table if exists (optional, for re-run safety)
-- DROP TABLE IF EXISTS customers;
-- DROP TABLE IF EXISTS accounts;

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

-- Create accounts table
-- CREATE TABLE accounts (
--     account_id    INT(11) NOT NULL AUTO_INCREMENT,
--     customer_id   INT(11) NOT NULL,
--     opening_date  DATETIME NOT NULL,
--     account_type  VARCHAR(10) NOT NULL,
--     amount        DECIMAL(10,2) NOT NULL DEFAULT 0.00,
--     status        TINYINT(1) NOT NULL,

--     -- Primary Key
--     PRIMARY KEY (account_id),

--     -- Foreign Key
--     CONSTRAINT fk_accounts_customer
--         FOREIGN KEY (customer_id)
--         REFERENCES customers(customer_id)
--         ON UPDATE CASCADE
--         ON DELETE RESTRICT
-- );

-- Insert sample data
-- INSERT INTO customers (customer_id, name, date_of_birth, city, zipcode, status) VALUES
-- (2000, 'Steve',  '1978-12-15', 'Delhi',          '110075', 1),
-- (2001, 'Arian',  '1988-05-21', 'Newburgh, NY',   '12550',  1),
-- (2002, 'Hadley', '1988-04-30', 'Englewood, NJ',  '07631',  1),
-- (2003, 'Ben',    '1988-01-04', 'Manchester, NH','03102',  0),
-- (2004, 'Nina',   '1988-05-14', 'Clarkston, MI',  '48348',  1),
-- (2005, 'Osman',  '1988-11-08', 'Hyattsville, MD','20782',  0);

-- INSERT INTO accounts
-- (account_id, customer_id, opening_date, account_type, amount, status)
-- VALUES
-- (95470, 2000, '2020-08-22 10:20:06', 'saving',   8823.23, 1),
-- (95471, 2002, '2020-08-09 10:27:22', 'checking', 3342.96, 1),
-- (95472, 2001, '2020-08-09 10:35:22', 'saving',   7000.00, 1),
-- (95473, 2001, '2020-08-09 10:38:22', 'saving',   5681.86, 1),
-- (95486, 2003, '2021-05-15 19:48:23', 'saving',   5000.00, 1);
