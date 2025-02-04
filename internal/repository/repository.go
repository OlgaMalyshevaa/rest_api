package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
)

type RepositoryInterface interface {
	Deposit(userID int, amount float64) error
	Transfer(fromUserID, toUserID int, amount float64) error
	GetLastTransactions(userID int, limit int) ([]Transaction, error)
}

type Repository struct {
	db *pgx.Conn
}

type Transaction struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	Amount    float64   `db:"amount"`
	Operation string    `db:"operation"`
	CreatedAt time.Time `db:"created_at"`
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Deposit(userID int, amount float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "UPDATE users SET balance = balance + $1 WHERE id = $2", amount, userID)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	_, err = tx.Exec(ctx, "INSERT INTO transactions (user_id, amount, operation) VALUES ($1, $2, 'deposit')", userID, amount)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}

func (r *Repository) Transfer(fromUserID, toUserID int, amount float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "UPDATE users SET balance = balance - $1 WHERE id = $2", amount, fromUserID)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	_, err = tx.Exec(ctx, "UPDATE users SET balance = balance + $1 WHERE id = $2", amount, toUserID)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	_, err = tx.Exec(ctx, "INSERT INTO transactions (user_id, amount, operation) VALUES ($1, $2, 'transfer')", fromUserID, -amount)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	_, err = tx.Exec(ctx, "INSERT INTO transactions (user_id, amount, operation) VALUES ($1, $2, 'transfer')", toUserID, amount)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}

func (r *Repository) GetLastTransactions(userID int, limit int) ([]Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.db.Query(ctx, "SELECT id, user_id, amount, operation, created_at FROM transactions WHERE user_id = $1 ORDER BY created_at DESC LIMIT $2", userID, limit)
	if err != nil {
		return nil,
			err
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var t Transaction
		if err := rows.Scan(&t.ID, &t.UserID, &t.Amount, &t.Operation, &t.CreatedAt); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}
