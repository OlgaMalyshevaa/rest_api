package repository

type MockRepository struct {
	DepositError          error
	TransferError         error
	LastTransactions      []Transaction
	LastTransactionsError error
}

func (m *MockRepository) Deposit(userID int, amount float64) error {
	return m.DepositError
}

func (m *MockRepository) Transfer(fromUserID, toUserID int, amount float64) error {
	return m.TransferError
}

func (m *MockRepository) GetLastTransactions(userID int, limit int) ([]Transaction, error) {
	return m.LastTransactions, m.LastTransactionsError
}
