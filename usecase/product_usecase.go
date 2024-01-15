package usecase

import (
	"Food-delivery/domain"
	"fmt"
)

type productUseCase struct {
	productRepo domain.ProductRepository
}

func NewProductUseCase(productRepo domain.ProductRepository) domain.ProductUseCase {
	return &productUseCase{productRepo: productRepo}
}

func (p *productUseCase) GetAll() ([]domain.Product, error) {

	pro, err := p.productRepo.GetAll()
	if err != nil {
		return nil, err
	}
	fmt.Println(pro)

	products := []domain.Product{}
	for _, product := range pro {
		products = append(products, domain.Product{
			Model:  product.Model,
			Name:   product.Name,
			Detail: product.Detail,
			Price:  product.Price,
		})
	}

	return products, nil
}

func (p *productUseCase) AddProduct(pro *domain.ProductForm) (*domain.ProductReply, error) {
	products := &domain.Product{
		Name:   pro.Name,
		Detail: pro.Detail,
		Price:  pro.Price,
	}
	product, err := p.productRepo.AddProduct(products)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *productUseCase) EditProduct(pro *domain.ProductForm, id uint) (*domain.ProductReply, error) {
	var product domain.Product
	_, err := p.productRepo.GetProductById(&product, id)
	if err != nil {
		return nil, err
	}

	goods := &domain.Product{
		Name:   pro.Name,
		Detail: pro.Detail,
		Price:  pro.Price,
	}

	products, err := p.productRepo.EditProduct(goods, id)

	if err != nil {
		return nil, err
	}

	return products, nil

}

func (p *productUseCase) DeleteProduct(id uint) error {
	var pro domain.Product
	err := p.productRepo.DeleteProduct(&pro, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *productUseCase) GetProductById(id uint) (*domain.ProductReply, error) {
	var req domain.Product
	product, err := p.productRepo.GetProductById(&req, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
