package repository
import (
        "gorm.io/gorm"

)

// ProductRepository handles database operations for the %s domain.
type ProductRepository struct {
        db *gorm.DB
}

// NewProductRepository creates a new Repository instance.
func NewProductRepository(db *gorm.DB) *ProductRepository {
return &ProductRepository{db: db}
}

// Add your repository methods here
