-- +goose Up
CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    amount DECIMAL (10, 2) NOT NULL,
    operation VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Insert
INSERT INTO transactions (user_id, amount, operation) VALUES
(1, 150.00, 'deposit'),
(1, 50.00, 'withdraw'),
(2, 200.00, 'deposit'),
(3, 100.00, 'transfer'),
(4, 20.00, 'transfer'),
(4, 150.00, 'deposit'),
(4, 50.00, 'withdraw'),
(5, 200.00, 'deposit'),
(5, 100.00, 'transfer'),
(6, 20.00, 'transfer');

-- +goose Down
DROP TABLE IF EXISTS transactions;
