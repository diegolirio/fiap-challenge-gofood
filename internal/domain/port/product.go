package port

import (
	"github.com/fiap/challenge-gofood/internal/domain/entity"
)

// Primary ports to Customer

type ProductUseCasePort interface {
	CreateProduct(name string, price float64, categoryID int) (*entity.Product, error)
	GetProductById(id uint) (*entity.Product, error)
	GetProducts() ([]*entity.Product, error)
	UpdateProduct(product *entity.Product) (*entity.Product, error)
	DeleteProduct(id uint) error
}

// Secondary ports to Product

type ProductRepositoryPort interface {
	CreateProduct(name string, price float64, categoryID int) (*entity.Product, error)
	GetProductById(id uint) (*entity.Product, error)
	GetProducts() ([]*entity.Product, error)
	UpdateProduct(product *entity.Product) (*entity.Product, error)
	DeleteProduct(id uint) error
}