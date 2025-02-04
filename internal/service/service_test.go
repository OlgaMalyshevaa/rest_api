package service

import (
	"errors"
	"file_rest_api/internal/repository"
	"testing"
)

func TestDeposit(t *testing.T) {
	mockRepo := &repository.MockRepository{}
	svc := NewService(mockRepo)

	err := svc.Deposit(1, 100)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	mockRepo.DepositError = errors.New("deposit error")
	err = svc.Deposit(1, 100)
	if err == nil {
		t.Errorf("expected error, got none")
	}
}

func TestTransfer(t *testing.T) {
	mockRepo := &repository.MockRepository{}
	svc := NewService(mockRepo)

	err := svc.Transfer(1, 2, 50)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	mockRepo.TransferError = errors.New("transfer error")
	err = svc.Transfer(1, 2, 50)
	if err == nil {
		t.Errorf("expected error, got none")
	}
}

func TestGetLastTransactions(t *testing.T) {
	mockRepo := &repository.MockRepository{}
	svc := NewService(mockRepo)

	mockRepo.LastTransactions = []repository.Transaction{
		{ID: 1, UserID: 1, Amount: 100.0, Operation: "deposit"},
		{ID: 2, UserID: 1, Amount: 50.0, Operation: "transfer"},
	}

	transactions, err := svc.GetLastTransactions(1, 10)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(transactions) != 2 {
		t.Errorf("expected 2 transactions, got %d", len(transactions))
	}

	mockRepo.LastTransactionsError = errors.New("error retrieving transactions")
	_, err = svc.GetLastTransactions(1, 10)
	if err == nil {
		t.Errorf("expected error, got none")
	}
}
