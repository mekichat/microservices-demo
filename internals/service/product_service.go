package service

import (
     "microservices-demo/internals/models"
	 "microservices-demo/internals/repository"
)


type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
    return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(name string, price int) models.Product {
	 product := models.Product{Name: name, Price: price}
     return s.repo.Add(product)
}

func (s *ProductService) ListProducts() []models.Product {
    return s.repo.GetAll()
}

func (s *ProductService) UpdateProduct(id uint, name string, price int) (models.Product, bool) {
    return s.repo.Update(id, name, price)
}

func (s *ProductService) DeleteProduct(id uint) bool {
    return s.repo.Delete(id)
}