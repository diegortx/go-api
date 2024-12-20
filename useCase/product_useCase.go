package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	productoId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}
	product.ID = productoId
	return product, nil
}

func (pu *ProductUseCase) GetProductById(id_product int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(id_product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pu *ProductUseCase) UpdateProduct(id_product int, product model.Product) (*model.Product, error) {
	productUpdated, err := pu.repository.UpdateProduct(id_product, product)
	if err != nil {
		return nil, err
	}
	return productUpdated, nil
}

func (pu *ProductUseCase) DeleteProductById(id_product int) (*model.Product, error) {
	product, err := pu.repository.DeleteProductById(id_product)

	if err != nil {
		return nil, err
	}
	return product, nil
}
