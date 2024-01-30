package repository

import (
	"context"

	"github.com/aldisypu/go-pos/internal/domain"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		DB: db,
	}
}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, transaction *domain.Transaction) error {
	return r.DB.WithContext(ctx).Create(transaction).Error
}

func (r *TransactionRepository) GetAllTransactions(ctx context.Context) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := r.DB.WithContext(ctx).Find(&transactions).Error
	return transactions, err
}

func (r *TransactionRepository) GetTransactionByID(ctx context.Context, id uint) (*domain.Transaction, error) {
	var transaction domain.Transaction
	err := r.DB.WithContext(ctx).First(&transaction, id).Error
	return &transaction, err
}

func (r *TransactionRepository) UpdateTransaction(ctx context.Context, id uint, updatedTransaction *domain.Transaction) error {
	return r.DB.WithContext(ctx).Model(&domain.Transaction{}).Where("id = ?", id).Updates(updatedTransaction).Error
}

func (r *TransactionRepository) DeleteTransaction(ctx context.Context, id uint) error {
	return r.DB.WithContext(ctx).Delete(&domain.Transaction{}, id).Error
}
