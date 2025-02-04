-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    balance DECIMAL (10, 2) DEFAULT 0.00
);

-- Insert
INSERT INTO users (name, balance) VALUES
('Natasha', 1000.0),
('Pasha', 500.0),
('Sasha', 100.0),
('Sveta', 150.0),
('Olya', 1500.0),
('Katya', 500.0),
('Alice', 100.0),
('Bob', 150.0);


-- +goose Down
DROP TABLE users;

