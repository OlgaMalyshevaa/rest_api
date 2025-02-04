package service

import "file_rest_api/internal/repository"

type Service struct {
	repo repository.RepositoryInterface
}

func NewService(repo repository.RepositoryInterface) *Service {
	return &Service{repo: repo}
}

func (s *Service) Deposit(userID int, amount float64) error {
	return s.repo.Deposit(userID, amount)
}

func (s *Service) Transfer(fromUserID, toUserID int, amount float64) error {
	return s.repo.Transfer(fromUserID, toUserID, amount)
}

func (s *Service) GetLastTransactions(userID int, limit int) ([]repository.Transaction, error) {
	return s.repo.GetLastTransactions(userID, limit)
}
