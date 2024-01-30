package service

import (
	"context"

	"github.com/aldisypu/go-pos/internal/domain"
	"github.com/aldisypu/go-pos/internal/repository"
)

type ProductService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, product *domain.Product) error {
	return s.productRepo.CreateProduct(ctx, product)
}

func (s *ProductService) GetAllProducts(ctx context.Context) ([]domain.Product, error) {
	return s.productRepo.GetAllProducts(ctx)
}

func (s *ProductService) GetProductByID(ctx context.Context, id uint) (*domain.Product, error) {
	return s.productRepo.GetProductByID(ctx, id)
}

func (s *ProductService) UpdateProduct(ctx context.Context, id uint, updatedProduct *domain.Product) error {
	return s.productRepo.UpdateProduct(ctx, id, updatedProduct)
}

func (s *ProductService) DeleteProduct(ctx context.Context, id uint) error {
	return s.productRepo.DeleteProduct(ctx, id)
}
