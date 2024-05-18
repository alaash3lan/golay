package handler

import (
"golay/internal/domain/product/service"

)

// productHandler handles HTTP requests related to product.
type ProductHandler struct {
service  *service.ProductService
}

// NewsHandler creates a new ProductHandler instance.
func NewProductHandler(service *service.ProductService) *ProductHandler {
return &ProductHandler{service: service}
}
// Add your handler methods here
