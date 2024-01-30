package usecase

import (
	"Food-delivery/domain"
)

type productUseCase struct {
	productRepo domain.ProductRepository
}

func NewProductUseCase(productRepo domain.ProductRepository) domain.ProductUseCase {
	return &productUseCase{productRepo: productRepo}
}

func (uc *productUseCase) GetAll() ([]domain.Product, error) {
	products, err := uc.productRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (uc *productUseCase) AddProduct(form *domain.ProductForm) (*domain.ProductReply, error) {
	products := &domain.Product{
		Name:   form.Name,
		Detail: form.Detail,
		Price:  form.Price,
	}

	product, err := uc.productRepo.Create(products)
	if err != nil {
		return nil, err
	}

	productReply := &domain.ProductReply{
		Name:   product.Name,
		Detail: product.Detail,
		Price:  product.Price,
	}

	return productReply, nil
}

func (uc *productUseCase) EditProduct(form *domain.ProductForm, id uint) (*domain.ProductReply, error) {
	var product domain.Product
	product.ID = id
	_, err := uc.productRepo.GetOneByID(&product)
	if err != nil {
		return nil, err
	}

	goods := &domain.Product{
		Name:   form.Name,
		Detail: form.Detail,
		Price:  form.Price,
	}
	goods.ID = id

	products, err := uc.productRepo.Edit(goods)
	if err != nil {
		return nil, err
	}

	productReply := &domain.ProductReply{
		Name:   products.Name,
		Detail: products.Detail,
		Price:  products.Price,
	}

	return productReply, nil
}

func (uc *productUseCase) DeleteProduct(id uint) error {
	var product domain.Product
	product.ID = id
	err := uc.productRepo.Delete(&product)
	if err != nil {
		return err
	}

	return nil
}

func (uc *productUseCase) GetProductById(id uint) (*domain.ProductReply, error) {
	var product domain.Product
	product.ID = id
	products, err := uc.productRepo.GetOneByID(&product)
	if err != nil {
		return nil, err
	}

	productReply := &domain.ProductReply{
		Name:   products.Name,
		Detail: products.Detail,
		Price:  products.Price,
	}

	return productReply, nil
}
