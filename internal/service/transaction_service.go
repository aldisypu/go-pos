package service

import (
	"context"

	"github.com/aldisypu/go-pos/internal/domain"
	"github.com/aldisypu/go-pos/internal/repository"
)

type TransactionService struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
	}
}

func (s *TransactionService) CreateTransaction(ctx context.Context, transaction *domain.Transaction) error {
	return s.transactionRepo.CreateTransaction(ctx, transaction)
}

func (s *TransactionService) GetAllTransactions(ctx context.Context) ([]domain.Transaction, error) {
	return s.transactionRepo.GetAllTransactions(ctx)
}

func (s *TransactionService) GetTransactionByID(ctx context.Context, id uint) (*domain.Transaction, error) {
	return s.transactionRepo.GetTransactionByID(ctx, id)
}

func (s *TransactionService) UpdateTransaction(ctx context.Context, id uint, updatedTransaction *domain.Transaction) error {
	return s.transactionRepo.UpdateTransaction(ctx, id, updatedTransaction)
}

func (s *TransactionService) DeleteTransaction(ctx context.Context, id uint) error {
	return s.transactionRepo.DeleteTransaction(ctx, id)
}
