package service

import (
	"context"
	"github.com/r1nb0/food-app/product-svc/internal/domain/models"
)

type CategoryService interface {
	Create(ctx context.Context, category models.CategoryCreate) (int64, error)
	Update(ctx context.Context, category models.Category) error
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]models.Category, error)
	GetByID(ctx context.Context, id int64) (models.Category, error)
}

type ProductService interface {
	Create(ctx context.Context, product models.ProductCreate) (int64, error)
	Update(ctx context.Context, product models.Product) error
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]models.Product, error)
	GetByID(ctx context.Context, id int64) (models.Product, error)
	GetByCategory(ctx context.Context, categoryID int64) ([]models.Product, error)
	GetDailyRecs(ctx context.Context) ([]models.Product, error)
}
