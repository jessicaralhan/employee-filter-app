-- Create the employees table
CREATE TABLE IF NOT EXISTS employees (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    status VARCHAR(50),
    country VARCHAR(100)
);

-- Insert sample data
INSERT INTO employees (name, status, country) VALUES 
('John Doe', 'active', 'USA'),
('Jane Smith', 'active', 'Canada'),
('Bob Wilson', 'inactive', 'UK'); 