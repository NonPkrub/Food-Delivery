package usecase

import (
	"Food-delivery/domain"
	"fmt"
)

type basketProductUseCase struct {
	basketProductRepo domain.BasketProductRepository
	promotionRepo     domain.PromotionRepository
	productRepo       domain.ProductRepository
	basketRepo        domain.BasketRepository
}

func NewBasketProductUseCase(basketProductRepo domain.BasketProductRepository, promotionRepo domain.PromotionRepository,
	productRepo domain.ProductRepository, basketRepo domain.BasketRepository) domain.BasketProductUseCase {
	return &basketProductUseCase{basketProductRepo: basketProductRepo, promotionRepo: promotionRepo, productRepo: productRepo, basketRepo: basketRepo}
}

func (uc *basketProductUseCase) AddProductInBasket(form *domain.BasketProduct, id uint) error {
	basket := &domain.Basket{
		UserID: id,
	}

	userBasket, err := uc.basketRepo.GetOneByID(basket)
	if err != nil {
		return err
	}

	baskets := &domain.BasketProduct{
		BasketID:  userBasket.ID,
		ProductID: form.ProductID,
		Quantity:  form.Quantity,
	}

	err = uc.basketProductRepo.Create(baskets)
	if err != nil {
		return err
	}

	return nil
}

func (uc *basketProductUseCase) EditProductInBasket(form *domain.BasketProduct, id uint) error {
	basket := &domain.Basket{
		UserID: id,
	}

	userBasket, err := uc.basketRepo.GetOneByID(basket)
	if err != nil {
		return err
	}

	baskets := &domain.BasketProduct{
		BasketID:  userBasket.ID,
		ProductID: form.ProductID,
		Quantity:  form.Quantity,
	}

	err = uc.basketProductRepo.Edit(baskets)
	if err != nil {
		return err
	}

	return nil
}

func (uc *basketProductUseCase) DeleteProductInBasket(form *domain.BasketProduct, id uint) error {
	basket := &domain.Basket{
		UserID: id,
	}

	userBasket, err := uc.basketRepo.GetOneByID(basket)
	if err != nil {
		return err
	}

	baskets := &domain.BasketProduct{
		BasketID:  userBasket.ID,
		ProductID: form.ProductID,
	}

	err = uc.basketProductRepo.Delete(baskets)
	if err != nil {
		return err
	}

	return nil
}

func (uc *basketProductUseCase) GetProductInBasket(form *domain.BasketProduct) ([]domain.BasketProductReply, float64, error) {
	basket := &domain.BasketProduct{
		BasketID:  form.BasketID,
		ProductID: form.ProductID,
		Quantity:  form.Quantity,
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
			Product: append(pb, domain.ProductForm{
				Name:   productDetail.Name,
				Detail: productDetail.Detail,
				Price:  productDetail.Price,
			}),
			Quantity: pro.Quantity,
		})

		totalPrice := calculateTotalPrice(productDetail.Price, pro.Quantity)
		totalProductPrice += totalPrice

		if promotionId != 0 && !oneTimeUse {
			promotion := &domain.PromotionProduct{}
			promotion.PromotionID = promotionId
			promotion.ProductID = pro.ProductID
			promotions, err := uc.promotionRepo.FindOneByID(promotion)
			if err != nil {
				return nil, 0, err
			}

			promotionDetail := &domain.Promotion{}
			promotionDetail.ID = promotions.PromotionID
			promotionDiscount, err := uc.promotionRepo.FindOne(promotionDetail)
			if err != nil {
				return nil, 0, err
			}

			fmt.Println(pro.ProductID, promotions.ProductID, pro.ProductID == promotions.ProductID)

			if pro.ProductID == promotions.ProductID {
				totalProductPrice = totalProductPrice - promotionDiscount.Discount
				oneTimeUse = true
			}

		}
	}

	totalPrices := float64(0)
	totalPrices = totalProductPrice

	return products, totalPrices, nil
}

func calculateTotalPrice(products float64, number uint) float64 {
	totalPrice := float64(0)
	totalPrice = products * float64(number)

	return totalPrice
}
