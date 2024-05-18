package service

import (
        "golay/internal/domain/product/repository"
)

// Service provides business logic for the product domain.
type ProductService struct {
        repo *repository.ProductRepository
}

// NewService creates a new Service instance.
func NewProductService(repo *repository.ProductRepository) *ProductService {
        return &ProductService{repo: repo}
}

// Add your service methods here
