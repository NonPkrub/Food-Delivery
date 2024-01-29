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

func (b *basketProductUseCase) GetProductInBasket(req *domain.BasketProduct) ([]domain.BasketProductReply, float64, error) {
	basket := &domain.BasketProduct{
		BasketID:  req.BasketID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	product, err := b.basketProductRepo.GetProductInBasket(basket)
	if err != nil {
		return nil, 0, err
	}

	var totalProductPrice float64

	pb := []domain.ProductForm{}

	promotionId, err := b.basketProductRepo.GetPromotionByBasketId(basket)
	if err != nil {
		return nil, 0, err
	}

	products := []domain.BasketProductReply{}
	for _, pro := range product {
		price, err := b.basketProductRepo.GetProductById(req, pro.ProductID)
		if err != nil {
			return nil, 0, err
		}

		if promotionId != 0 {
			promotion, err := b.basketProductRepo.GetPromotionBasket(basket, promotionId)
			if err != nil {
				return nil, 0, err
			}
			fmt.Println(pro.ProductID == promotion.ProductID, pro.ProductID)
			if pro.ProductID == promotion.ProductID {
				price.Price = price.Price - promotion.Discount
			}

		}

		products = append(products, domain.BasketProductReply{
			BasketID: pro.BasketID,
			Product: append(pb, domain.ProductForm{
				Name:   price.Name,
				Detail: price.Detail,
				Price:  price.Price,
			}),
			Quantity: pro.Quantity,
		})

		totalPrice := calculateTotalPrice(price.Price, pro.Quantity)
		totalProductPrice += totalPrice
	}

	var totalPrices float64
	totalPrices = totalProductPrice

	// fmt.Println(promotion)
	// if promotion.ProductID != 0 {
	// 	totalPrices = totalProductPrice - promotion.Discount
	// }
	// if promotionId != 0 {
	// 	promotion, err := b.basketProductRepo.GetPromotionBasket(basket, promotionId)
	// 	if err != nil {
	// 		return nil, 0, err
	// 	}

	// 	if products.ProductID == promotion.ProductID {
	// 		totalPrices = totalProductPrice - promotion.Discount
	// 	}

	// }
	// fmt.Println(totalPrices)

	return products, totalPrices, nil

}

func calculateTotalPrice(products float64, number uint) float64 {
	var totalPrice float64
	totalPrice = products * float64(number)

	return totalPrice
}
