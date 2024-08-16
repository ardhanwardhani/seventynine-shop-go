package application

import (
	"seventynine-shop-go/internal/domain"
)

type ProductService struct {
	repository domain.ProductRepository
}

func NewProductService(repo domain.ProductRepository) *ProductService {
	return &ProductService{repository: repo}
}

func (s *ProductService) GetAllProducts() ([]domain.Product, error) {
	return s.repository.GetAll()
}

func (s *ProductService) GetProductByID(id int) (*domain.Product, error) {
	return s.repository.GetByID(id)
}

func (s *ProductService) CreateProduct(product *domain.Product) error {
	return s.repository.Create(product)
}

func (s *ProductService) UpdateProduct(product *domain.Product) error {
	return s.repository.Update(product)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.repository.Delete(id)
}
