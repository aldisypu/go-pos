package repository

import (
	"context"

	"github.com/aldisypu/go-pos/internal/domain"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (r *ProductRepository) CreateProduct(ctx context.Context, product *domain.Product) error {
	return r.DB.WithContext(ctx).Create(product).Error
}

func (r *ProductRepository) GetAllProducts(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product
	err := r.DB.WithContext(ctx).Find(&products).Error
	return products, err
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id uint) (*domain.Product, error) {
	var product domain.Product
	err := r.DB.WithContext(ctx).First(&product, id).Error
	return &product, err
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, id uint, updatedProduct *domain.Product) error {
	return r.DB.WithContext(ctx).Model(&domain.Product{}).Where("id = ?", id).Updates(updatedProduct).Error
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, id uint) error {
	return r.DB.WithContext(ctx).Delete(&domain.Product{}, id).Error
}
