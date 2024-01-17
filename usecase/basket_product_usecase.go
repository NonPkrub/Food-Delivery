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

func (b *basketProductUseCase) AddProductInBasket(req *domain.BasketProduct) error {
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

func (b *basketProductUseCase) EditProductInBasket(req *domain.BasketProduct) error {
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

func (b *basketProductUseCase) DeleteProductInBasket(req *domain.BasketProduct) error {
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

func (b *basketProductUseCase) GetProductInBasket(req *domain.BasketProduct) ([]domain.BasketProductReply, error) {
	basket := &domain.BasketProduct{
		BasketID:  req.BasketID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	product, err := b.basketProductRepo.GetProductInBasket(basket)
	if err != nil {
		return nil, err
	}

	var totalProductPrice float64

	pb := []domain.Product{}

	products := []domain.BasketProductReply{}
	for _, pro := range product {
		price, err := b.basketProductRepo.GetProductById(req, pro.ProductID)
		if err != nil {
			return nil, err
		}

		totalPrice := calculateTotalPrice(price.Price, pro.Quantity)
		fmt.Println(totalPrice)
		totalProductPrice += totalPrice
		fmt.Println(totalProductPrice)

		products = append(products, domain.BasketProductReply{
			BasketID: pro.BasketID,
			Product: append(pb, domain.Product{
				Name:   price.Name,
				Detail: price.Detail,
				Price:  price.Price,
			}),
			Quantity: pro.Quantity,
		})

	}

	return products, nil

}

func calculateTotalPrice(products float64, number uint) float64 {
	var totalPrice float64
	totalPrice = products * float64(number)

	return totalPrice
}
