package usecase

import (
	"Food-delivery/domain"
	"fmt"
)

type basketProductUseCase struct {
	basketProductRepo domain.BasketProductRepository
	promotionRepo     domain.PromotionRepository
	productRepo       domain.ProductRepository
}

func NewBasketProductUseCase(basketProductRepo domain.BasketProductRepository, promotionRepo domain.PromotionRepository,
	productRepo domain.ProductRepository) domain.BasketProductUseCase {
	return &basketProductUseCase{basketProductRepo: basketProductRepo, promotionRepo: promotionRepo, productRepo: productRepo}
}

func (uc *basketProductUseCase) AddProductInBasket(req *domain.BasketProduct) error {
	basket := &domain.BasketProduct{
		BasketID:  req.BasketID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	err := uc.basketProductRepo.Create(basket)
	if err != nil {
		return err
	}

	return nil
}

func (uc *basketProductUseCase) EditProductInBasket(req *domain.BasketProduct) error {
	basket := &domain.BasketProduct{
		BasketID:  req.BasketID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	err := uc.basketProductRepo.Edit(basket)
	if err != nil {
		return err
	}

	return nil
}

func (uc *basketProductUseCase) DeleteProductInBasket(req *domain.BasketProduct) error {
	basket := &domain.BasketProduct{
		BasketID:  req.BasketID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	err := uc.basketProductRepo.Delete(basket)
	if err != nil {
		return err
	}

	return nil
}

func (uc *basketProductUseCase) GetProductInBasket(req *domain.BasketProduct) ([]domain.BasketProductReply, float64, error) {
	basket := &domain.BasketProduct{
		BasketID:  req.BasketID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	product, err := uc.basketProductRepo.FindAllByID(basket)
	if err != nil {
		return nil, 0, err
	}

	var totalProductPrice float64

	pb := []domain.ProductForm{}

	promotionId, err := uc.basketProductRepo.GetPromotionByBasketID(basket)
	if err != nil {
		return nil, 0, err
	}

	products := []domain.BasketProductReply{}
	oneTimeUse := false

	for _, pro := range product {
		productID := &domain.Product{}
		productID.ID = pro.ProductID
		productDetail, err := uc.productRepo.GetOneByID(productID)
		if err != nil {
			return nil, 0, err
		}

		// if promotionId != 0 {
		// 	promotion, err := b.basketProductRepo.GetPromotionBasket(basket, promotionId)
		// 	if err != nil {
		// 		return nil, 0, err
		// 	}
		// 	fmt.Println(pro.ProductID == promotion.ProductID, pro.ProductID)
		// 	if pro.ProductID == promotion.ProductID {
		// 		price.Price = price.Price - promotion.Discount
		// 	}
		// }

		products = append(products, domain.BasketProductReply{
			BasketID: pro.BasketID,
			Product: append(pb, domain.ProductForm{
				Name:   productDetail.Name,
				Detail: productDetail.Detail,
				Price:  productDetail.Price,
			}),
			Quantity: pro.Quantity,
		})

		totalPrice := calculateTotalPrice(productDetail.Price, pro.Quantity)
		totalProductPrice += totalPrice

		fmt.Println(!oneTimeUse)
		if promotionId != 0 && !oneTimeUse {
			promotion := &domain.PromotionProduct{}
			promotion.PromotionID = promotionId
			promotions, err := uc.promotionRepo.GetOneByID(promotion)
			if err != nil {
				return nil, 0, err
			}

			promotionDetail := &domain.Promotion{}
			promotionDetail.ID = promotions.PromotionID
			promotionDiscount, err := uc.promotionRepo.FindOne(promotionDetail)
			if err != nil {
				return nil, 0, err
			}

			if pro.ProductID == promotions.ProductID {
				totalProductPrice = totalProductPrice - promotionDiscount.Discount
			}

			oneTimeUse = true

		}
	}

	var totalPrices float64
	totalPrices = totalProductPrice

	return products, totalPrices, nil
}

func calculateTotalPrice(products float64, number uint) float64 {
	var totalPrice float64
	totalPrice = products * float64(number)

	return totalPrice
}
