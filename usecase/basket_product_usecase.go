package usecase

import (
	"Food-delivery/domain"
	"fmt"
)

type basketProductUseCase struct {
	basketProductRepo domain.BasketProductRepository
}

func NewBasketProductUseCase(basketProductRepo domain.BasketProductRepository) domain.BasketProductUseCase {
	return &basketProductUseCase{basketProductRepo: basketProductRepo}
}

func (b *basketProductUseCase) AddProductInBasket(req *domain.BasketProductForm) error {
	basket := &domain.BasketProduct{
		BasketID:  req.BasketID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	err := b.basketProductRepo.AddProductInBasket(basket)
	if err != nil {
		return err
	}

	return nil
}

func (b *basketProductUseCase) EditProductInBasket(req *domain.BasketProductForm) error {
	basket := &domain.BasketProduct{
		BasketID:  req.BasketID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	err := b.basketProductRepo.EditProductInBasket(basket)
	if err != nil {
		return err
	}

	return nil
}

func (b *basketProductUseCase) DeleteProductInBasket(req *domain.BasketProductForm) error {
	basket := &domain.BasketProduct{
		BasketID:  req.BasketID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	err := b.basketProductRepo.DeleteProductInBasket(basket)
	if err != nil {
		return err
	}

	return nil
}

func (b *basketProductUseCase) GetProductInBasket(req *domain.BasketProductForm) ([]domain.BasketProductReply, error) {
	basket := &domain.BasketProduct{
		BasketID:  req.BasketID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	product, err := b.basketProductRepo.GetProductInBasket(basket)
	if err != nil {
		return nil, err
	}
	fmt.Println(product)

	products := []domain.BasketProductReply{}
	for _, pro := range product {
		products = append(products, domain.BasketProductReply{
			BasketID:  pro.BasketID,
			ProductID: pro.ProductID,
			Quantity:  pro.Quantity,
		})

		price, err := b.basketProductRepo.GetProductById(req, pro.ProductID)
		if err != nil {
			return nil, err
		}
		fmt.Println(price.Product)
	}

	return products, nil

}
