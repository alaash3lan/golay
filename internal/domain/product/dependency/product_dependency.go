package dependency
import (
        "golay/internal/domain/product/handler"
        "golay/internal/domain/product/repository"
        "golay/internal/domain/product/service"

        "gorm.io/gorm"
)

func SetupProductDependencies(db *gorm.DB) (*handler.ProductHandler, error) {
        productRepo := repository.NewProductRepository(db)

        productService := service.NewProductService(productRepo)

        productHandler := handler.NewProductHandler(productService)

        return productHandler, nil
}
